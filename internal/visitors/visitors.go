package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/scip-code/scip-go/internal/document"
	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

// OriginFile returns the `//line`-adjusted source path for pos, cleaned.
//
// cgo (and any generated code carrying `//line` directives) is compiled from
// files that the go command rewrites into the build cache -- e.g. a cgo file
// `foo.go` becomes `foo.cgo1.go` under GOCACHE. `Fset.File(pos).Name()` returns
// that physical cache path, but `Fset.Position(pos)` honors the `//line`
// directives cgo emits and resolves back to the real `.go` source. Keying
// documents by this origin keeps occurrences anchored to source the repo
// actually contains, instead of an ephemeral cache path. For ordinary files the
// origin is the file itself, so non-generated packages are unaffected.
func OriginFile(pkg *packages.Package, pos token.Pos) string {
	return CleanResolve(pkg.Fset.Position(pos).Filename)
}

// CleanResolve returns path with symlinks resolved and cleaned. Resolving
// symlinks keeps `//line`-derived origins comparable to pkg.GoFiles even when a
// module is reached through a symlinked directory (otherwise a real source file
// could be dropped as "not a GoFile"). Falls back to Clean when the path can't
// be resolved -- e.g. a `//line` target that names a file not on disk (a yacc
// `.y` grammar, or a build-cache path) -- which then simply won't match any
// GoFile and is dropped, as intended.
func CleanResolve(path string) string {
	if resolved, err := filepath.EvalSymlinks(path); err == nil {
		return resolved
	}
	return filepath.Clean(path)
}

// RealGoFiles is the set of a package's on-disk source files (resolved paths).
// An occurrence whose origin is not in this set comes from generated glue with
// no real source (cgo's `_cgo_gotypes.go`, a yacc `.y`, compiler-inserted thunks
// lacking a `//line`); it is dropped rather than mis-attributed.
func RealGoFiles(pkg *packages.Package) map[string]struct{} {
	set := make(map[string]struct{}, len(pkg.GoFiles))
	for _, f := range pkg.GoFiles {
		set[CleanResolve(f)] = struct{}{}
	}
	return set
}

func VisitPackageSyntax(
	moduleRoot string,
	pkg *packages.Package,
	pathToDocuments map[string]*document.Document,
	globalSymbols *lookup.Global,
) {
	pkgSymbols := lookup.NewPackageSymbols(pkg)
	goFiles := RealGoFiles(pkg)
	// Iterate over all the files, collect any global symbols
	for _, f := range pkg.Syntax {

		origin := OriginFile(pkg, f.Package)
		relative, _ := filepath.Rel(moduleRoot, origin)

		// Always visit to collect package-level symbols, but only keep a
		// document for files that map to real source. Generated files (e.g.
		// cgo's `_cgo_gotypes.go`) resolve to a non-source origin; their
		// occurrences are compiler glue and must not become a document.
		doc := visitSyntax(pkg, pkgSymbols, f, relative, origin)
		if _, ok := goFiles[origin]; ok {
			pathToDocuments[origin] = doc
		}
	}

	globalSymbols.Add(pkgSymbols)
}

func visitSyntax(pkg *packages.Package, pkgSymbols *lookup.Package, f *ast.File, relative, originAbs string) *document.Document {
	doc := document.NewDocument(relative, originAbs, pkg, pkgSymbols)

	// TODO: Maybe we should do this before? we have traverse all
	// the fields first before, but now I think it's fine right here
	// .... maybe
	visitTypesInFile(doc, pkg, f)

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.BadDecl:
			continue

		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				// These do not create global symbols
				continue

			case token.TYPE:
				// We do this via visitTypesInFile above

			case token.VAR, token.CONST:
				// visit var
				visitVarDefinition(doc, pkg, decl)

			default:
				panic("Unhandled general declaration")
			}

		case *ast.FuncDecl:
			visitFunctionDefinition(doc, pkg, decl)
		}

	}

	return doc
}

func walkExprList(v ast.Visitor, list []ast.Expr) {
	for _, x := range list {
		ast.Walk(v, x)
	}
}

func walkDeclList(v ast.Visitor, list []ast.Decl) {
	for _, x := range list {
		ast.Walk(v, x)
	}
}

func descriptorTerm(name string) *scip.Descriptor {
	return &scip.Descriptor{
		Name:   name,
		Suffix: scip.Descriptor_Term,
	}
}

func scipRange(start, end token.Position, obj types.Object) scip.Range {
	var adjustment int32 = 0
	if pkgName, ok := obj.(*types.PkgName); ok && strings.HasPrefix(pkgName.Name(), `"`) {
		adjustment = 1
	}

	startLine := int32(start.Line - 1)
	startColumn := int32(start.Column - 1)
	endLine := int32(end.Line - 1)
	endColumn := int32(end.Column - 1)
	return scip.Range{
		Start: scip.Position{Line: startLine, Character: startColumn + adjustment},
		End:   scip.Position{Line: endLine, Character: endColumn - adjustment},
	}
}

func getIdentOfTypeExpr(pkg *packages.Package, ty ast.Expr) []*ast.Ident {
	switch ty := ty.(type) {
	case *ast.Ident:
		return []*ast.Ident{ty}
	case *ast.SelectorExpr:
		return []*ast.Ident{ty.Sel}
	case *ast.StarExpr:
		return getIdentOfTypeExpr(pkg, ty.X)
	case *ast.IndexExpr:
		return getIdentOfTypeExpr(pkg, ty.X)
	case *ast.BinaryExpr:
		// As far as I can tell, binary exprs are ONLY for type constraints
		// and those don't really define anything on the struct.
		//
		// So far now, we'll just not return anything.
		//
		// return append(s.getIdentOfTypeExpr(ty.X), s.getIdentOfTypeExpr(ty.Y)...)
		return []*ast.Ident{}
	case *ast.UnaryExpr:
		return getIdentOfTypeExpr(pkg, ty.X)

	case *ast.IndexListExpr:
		return getIdentOfTypeExpr(pkg, ty.X)

	// TODO: Should see if any of these need better ident finders
	case *ast.InterfaceType:
		return nil
	case *ast.FuncType:
		return nil
	case *ast.FuncLit:
		return nil
	case *ast.MapType:
		return nil
	case *ast.ArrayType:
		return nil
	case *ast.ChanType:
		return nil

	default:
		slog.Debug(fmt.Sprintf("Unhandled named struct field: %T %+v\n%s", ty, ty, pkg.Fset.Position(ty.Pos())))
		return nil
	}
}
