package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log/slog"

	"github.com/scip-code/scip-go/internal/document"
	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip-go/internal/newtypes"
	"github.com/scip-code/scip-go/internal/symbols"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func NewFileVisitor(
	doc *document.Document,
	pkg *packages.Package,
	file *ast.File,
	pkgSymbols *lookup.Package,
	globalSymbols *lookup.Global,
	docs map[string]*document.Document,
) *fileVisitor {
	caseClauses := map[token.Pos]types.Object{}
	for implicit, obj := range pkg.TypesInfo.Implicits {
		if _, ok := implicit.(*ast.CaseClause); ok {
			caseClauses[obj.Pos()] = obj
		}
	}

	return &fileVisitor{
		doc:           doc,
		docs:          docs,
		pkg:           pkg,
		file:          file,
		locals:        map[token.Pos]lookup.Local{},
		pkgSymbols:    pkgSymbols,
		globalSymbols: globalSymbols,
		caseClauses:   caseClauses,
	}
}

// fileVisitor visits an entire file, but it must be called
// after StructVisitor.
//
// Iterates over a file,
type fileVisitor struct {
	// doc is the document for this file's own origin; SymbolInformation defined in
	// the file is attached here in Finish. Occurrences are routed per-occurrence
	// to the document of their //line-adjusted origin (see docs) -- usually doc,
	// but a generated file (e.g. cgo's rewritten source) can attribute some of its
	// occurrences to a different real source file.
	doc *document.Document

	// docs maps a resolved-origin absolute path to its document (the same map the
	// index builds -- one entry per real source file). Occurrences are routed here
	// by origin; an occurrence whose origin has no entry (cgo glue, a yacc `.y`, a
	// build-cache path) is dropped rather than mis-attributed.
	docs map[string]*document.Document

	// Current file information
	pkg  *packages.Package
	file *ast.File

	// local definition position to symbol and its type information
	locals map[token.Pos]lookup.Local

	// field definition position to symbol for the package
	pkgSymbols *lookup.Package

	// field definition position to symbol for the entire compliation
	globalSymbols *lookup.Global

	// caseClauses maps particular positions to different types for case clauses
	caseClauses map[token.Pos]types.Object

	// currentFuncDecl tracks the enclosing FuncDecl during Visit traversal
	currentFuncDecl *ast.FuncDecl
}

// Implements ast.Visitor
var _ ast.Visitor = &fileVisitor{}

func (v *fileVisitor) createNewLocalSymbol(pos token.Pos, obj types.Object) string {
	if _, ok := v.locals[pos]; ok {
		panic("Cannot create a new local symbol for an ident that has already been created")
	}

	symbol := fmt.Sprintf("local %d", len(v.locals))

	v.locals[pos] = lookup.Local{
		Symbol: symbol,
		Obj:    obj,
	}

	return symbol
}

func (v *fileVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.ImportSpec:
		// Use PkgNameOf to resolve the import reliably via the type checker.
		pkgName := v.pkg.TypesInfo.PkgNameOf(node)
		if pkgName == nil {
			slog.Warn("Could not find node", "node.Path", node.Path)
			return nil
		}

		importedPackage := v.pkg.Imports[pkgName.Imported().Path()]
		if importedPackage == nil {
			slog.Warn("Could not find node", "node.Path", node.Path)
			return nil
		}

		if node.Name != nil && node.Name.Name != "." && node.Name.Name != "_" {
			if sym, ok := v.globalSymbols.GetPkgSymbol(importedPackage); ok {
				namePos := v.pkg.Fset.Position(node.Name.Pos())
				v.newReference(namePos, sym, symbols.RangeFromName(
					namePos, node.Name.Name, false), false)
			}
		}

		position := v.pkg.Fset.Position(node.Path.Pos())
		v.emitImportReference(v.globalSymbols, position, importedPackage)

		return nil

	case *ast.SelectorExpr:
		if ident, ok := node.X.(*ast.Ident); ok {
			use := v.pkg.TypesInfo.Uses[ident]

			// We special case handling PkgNames because they do some goofy things
			// compared to almost every other construct in the language.
			switch sel := use.(type) {
			case *types.PkgName:
				startPosition := v.pkg.Fset.Position(ident.Pos())
				endPosition := v.pkg.Fset.Position(ident.End())

				packageID := newtypes.GetFromTypesPackage(sel.Imported())
				sym, ok := v.globalSymbols.GetPkgSymbolByID(packageID)
				if !ok {
					slog.Debug(fmt.Sprintf(
						"Missing symbol for package: %s", sel.Imported().Path()))
					return nil
				}

				symRange := scipRange(startPosition, endPosition, sel)
				v.newReference(startPosition, sym, symRange, false)

				// Then walk the selection
				ast.Walk(v, node.Sel)

				// and since we've handled the rest, end visit
				return nil
			}
		}

		return v
	case *ast.FuncDecl:
		v.currentFuncDecl = node
		if node.Doc != nil {
			ast.Walk(v, node.Doc)
		}
		if node.Recv != nil {
			ast.Walk(v, node.Recv)
		}
		ast.Walk(v, node.Name)
		ast.Walk(v, node.Type)
		if node.Body != nil {
			ast.Walk(v, node.Body)
		}
		return nil
	case *ast.File:
		if node.Doc != nil {
			ast.Walk(v, node.Doc)
		}

		// Handle package name declaration separately
		// No need to: ast.Walk(v, n.Name)

		walkDeclList(v, node.Decls)
		return nil
	case *ast.Ident:
		// Short circuit if this is a blank identifier
		if node.Name == "_" {
			return nil
		}

		startPosition := v.pkg.Fset.Position(node.Pos())
		endPosition := v.pkg.Fset.Position(node.End())

		// Short circuit on case clauses
		if obj, ok := v.caseClauses[node.Pos()]; ok {
			symName := v.createNewLocalSymbol(obj.Pos(), obj)
			v.newDefinition(startPosition, symName, scipRange(startPosition, endPosition, obj), nil, false)
			return nil
		}

		info := v.pkg.TypesInfo

		// Emit Definition
		def := info.Defs[node]
		if def != nil {
			var symName string
			if pkgSymbols, ok := v.pkgSymbols.GetSymbol(def.Pos()); ok {
				symName = pkgSymbols
			} else if globalSymbol, ok := v.globalSymbols.GetSymbol(v.pkg, def.Pos()); ok {
				symName = globalSymbol
			} else {
				symName = v.createNewLocalSymbol(def.Pos(), def)
			}

			v.newDefinition(
				startPosition,
				symName,
				scipRange(startPosition, endPosition, def),
				v.enclosingRange(node),
				v.isDeprecated(def.Pos()))
		}

		// Emit Reference
		ref := info.Uses[node]
		if ref != nil {
			if _, ok := ref.(*types.PkgName); ok {
				return v
			}

			var (
				symbol     string
				deprecated bool
			)

			if localSymbol, ok := v.locals[ref.Pos()]; ok {
				symbol = localSymbol.Symbol
			} else {
				var err error
				symInfo, ok, err := v.globalSymbols.GetSymbolOfObject(ref)
				if err != nil {
					slog.Debug(fmt.Sprintf(
						"Unable to find symbol of object: %s\nNode Position -> %s\n\nPath: %s\n\n",
						err,
						v.pkg.Fset.Position(node.Pos()),
						ref.Pkg().Path(),
					))
					return v
				}

				if !ok {
					return v
				}

				// Set the resulting info
				symbol = symInfo.Symbol
				deprecated = document.IsDocDeprecated(symInfo.Documentation)
			}

			v.newReference(startPosition, symbol, scipRange(startPosition, endPosition, ref), deprecated)
		}

		if def == nil && ref == nil {
			slog.Debug(fmt.Sprintf(
				"Neither def nor ref found: %s | %T | %s",
				node.Name,
				node,
				v.pkg.Fset.Position(node.Pos()),
			))
		}
	}

	return v
}

func (v *fileVisitor) emitImportReference(
	globalSymbols *lookup.Global,
	position token.Position,
	importedPackage *packages.Package,
) {
	sym, ok := globalSymbols.GetPkgSymbol(importedPackage)
	if !ok {
		slog.Debug(fmt.Sprintf("Missing symbol information for package: %s", importedPackage.ID))
		return
	}

	v.newReference(position, sym, symbols.RangeFromName(position, importedPackage.PkgPath, true), false)
}

// targetDoc resolves the document an occurrence at pos with range rng belongs
// to, along with the range to emit. A nil document means drop the occurrence.
//
// An occurrence's true home is its //line-adjusted origin file: usually the file
// being walked, but a generated file (cgo, ...) can attribute an occurrence to a
// *different* real source file, in which case it is routed there rather than
// dropped. An occurrence whose origin is not a real source document (cgo glue, a
// yacc `.y`, a build-cache path) is dropped.
//
// A range that escapes the origin's source gets one repair attempt first: cgo
// widens ranges by mangling identifiers (`C.puts` -> `_Cfunc_puts`), which pushes
// a correctly-positioned occurrence past end-of-line, and Document.RepairRange
// re-measures it against the real source. Only when that fails is the occurrence
// dropped, rather than emitted with a bogus location (which downstream SCIP
// consumers reject).
func (v *fileVisitor) targetDoc(pos token.Position, rng scip.Range) (*document.Document, scip.Range) {
	doc := v.docs[CleanResolve(pos.Filename)]
	if doc == nil {
		return nil, rng
	}
	if doc.InBounds(rng) {
		return doc, rng
	}
	if repaired, ok := doc.RepairRange(rng); ok {
		return doc, repaired
	}
	return nil, rng
}

// newDefinition emits a scip.Occurence ONLY. This will not emit a
// new symbol. You must do that using DeclareNewSymbol[ForPos]
func (v *fileVisitor) newDefinition(
	pos token.Position, symbol string, rng scip.Range, enclRng *scip.Range, deprecated bool,
) {
	doc, rng := v.targetDoc(pos, rng)
	if doc == nil {
		return
	}
	occ := &scip.Occurrence{
		TypedRange:  rng.AsTypedRange(),
		Symbol:      symbol,
		SymbolRoles: int32(scip.SymbolRole_Definition),
	}
	// Keep the enclosing range only if it fits the same source (a cgo-expanded
	// body can push it past EOF even when the name range is fine).
	if enclRng != nil && doc.InBounds(*enclRng) {
		occ.TypedEnclosingRange = enclRng.AsTypedEnclosingRange()
	}
	if deprecated {
		occ.Diagnostics = deprecatedDiagnostics()
	}
	doc.AppendOccurrence(occ)
}

func (v *fileVisitor) newReference(
	pos token.Position, symbol string, rng scip.Range, deprecated bool,
) {
	doc, rng := v.targetDoc(pos, rng)
	if doc == nil {
		return
	}
	occ := &scip.Occurrence{
		TypedRange:  rng.AsTypedRange(),
		Symbol:      symbol,
		SymbolRoles: int32(scip.SymbolRole_ReadAccess),
	}
	if deprecated {
		occ.Diagnostics = deprecatedDiagnostics()
	}
	doc.AppendOccurrence(occ)
}

// Finish attaches the file's SymbolInformation (package-level symbols defined in
// the file, plus locals) to the file's own document. Occurrences are routed to
// their origin documents during the walk; call Finish once after the walk.
func (v *fileVisitor) Finish() {
	documentFile := v.pkg.Fset.File(v.file.Pos())
	if documentFile == nil {
		return
	}

	documentSymbols := v.pkgSymbols.SymbolsForFile(documentFile)
	for _, local := range v.locals {
		symbolInfo := &scip.SymbolInformation{
			Symbol: local.Symbol,
		}

		if obj := local.Obj; obj != nil {
			symbolInfo.DisplayName = obj.Name()
			symbolInfo.Kind = symbols.KindForObject(obj)
			// Skip SignatureDocumentation for type-switch locals because
			// multiple case clauses share the same obj.Pos(), making the
			// type recorded in caseClauses nondeterministic.
			if _, isTypeSwitchLocal := v.caseClauses[obj.Pos()]; !isTypeSwitchLocal {
				if txt := local.SignatureText(); txt != "" {
					symbolInfo.SignatureDocumentation = &scip.Signature{
						Language: "go",
						Text:     txt,
					}
				}
			}
		}

		documentSymbols = append(documentSymbols, symbolInfo)
	}

	v.doc.AddSymbols(documentSymbols)
}

func (v *fileVisitor) enclosingRange(n *ast.Ident) *scip.Range {
	if v.currentFuncDecl == nil || v.currentFuncDecl.Name != n {
		return nil
	}
	startPosition := v.pkg.Fset.Position(v.currentFuncDecl.Pos())
	endPosition := v.pkg.Fset.Position(v.currentFuncDecl.End())
	rng := scipRange(startPosition, endPosition, v.pkg.TypesInfo.Defs[n])
	return &rng
}

func deprecatedDiagnostics() []*scip.Diagnostic {
	return []*scip.Diagnostic{{
		Severity: scip.Severity_Warning,
		Message:  "Deprecated",
		Tags:     []scip.DiagnosticTag{scip.DiagnosticTag_Deprecated},
	}}
}

// isDeprecated reports whether the symbol at pos has documentation containing
// a "Deprecated:" paragraph per Go convention.
func (v *fileVisitor) isDeprecated(pos token.Pos) bool {
	if info, ok := v.pkgSymbols.Get(pos); ok {
		return document.IsDocDeprecated(info.Documentation)
	}
	if info, ok := v.globalSymbols.GetSymbolInformation(v.pkg, pos); ok {
		return document.IsDocDeprecated(info.Documentation)
	}
	return false
}
