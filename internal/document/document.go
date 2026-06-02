package document

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/format"
	"go/token"
	"go/types"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip-go/internal/symbols"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

// indent is used to format struct fields.
const indent = "    "

// IsDocDeprecated reports whether the documentation strings contain a paragraph
// starting with "Deprecated:" per Go convention.
func IsDocDeprecated(docs []string) bool {
	for _, d := range docs {
		for line := range strings.SplitSeq(d, "\n") {
			if strings.HasPrefix(line, "Deprecated:") {
				return true
			}
		}
	}
	return false
}

func NewDocument(
	relative string,
	pkg *packages.Package,
	pkgSymbols *lookup.Package,
) *Document {
	return &Document{
		RelativePath: relative,
		pkg:          pkg,
		pkgSymbols:   pkgSymbols,

		docPkg: &doc.Package{},
	}
}

type Document struct {
	// Document relative path. To be used for scip.Document
	RelativePath string

	// The occurrence for `package foo` at the top of a Go file.
	//   It could be a definition or a reference, depending on the package structure.
	//   It doesn't get traversed in the same way as other parts of the tree,
	//   so we special case it here. It must get added to the occurences when
	//   creating a visitors.fileVisitor
	PackageOccurrence *scip.Occurrence

	// The package this document is contained in
	pkg *packages.Package

	// Hold information for docstrings and pretty linking
	docPkg *doc.Package

	// pkgSymbols maps positions to symbol names within
	// this document.
	pkgSymbols *lookup.Package
}

func (d *Document) GetSymbol(pos token.Pos) (string, bool) {
	return d.pkgSymbols.GetSymbol(pos)
}

// SetSymbolInformation registers a pre-built SymbolInformation at the given position.
func (d *Document) SetSymbolInformation(
	pos token.Pos, info *scip.SymbolInformation,
) {
	d.pkgSymbols.Set(pos, info)
}

// SetNewSymbol declares a new symbol and tracks it within a Document.
//
// NOTE: Does NOT emit a new occurrence
func (d *Document) SetNewSymbol(
	symbol string,
	parent ast.Node,
	ident *ast.Ident,
) {
	d.SetNewSymbolForPos(symbol, parent, ident, ident.Pos())
}

// SetNewSymbolForPos declares a new symbol and tracks it within a Document
// but allows for an override of the position. Generally speaking, you should use
// DeclareNewSymbol instead (since it will calculate the pos for most cases)
//
// NOTE: Does NOT emit a new occurrence
func (d *Document) SetNewSymbolForPos(
	symbol string,
	parent ast.Node,
	ident *ast.Ident,
	pos token.Pos,
) {
	var displayName string
	var documentation []string
	var sigDoc *scip.Signature
	var def types.Object

	if ident != nil {
		displayName = ident.Name

		def = d.pkg.TypesInfo.Defs[ident]
		if def != nil {
			if signature := typeStringForObject(def); signature != "" {
				sigDoc = &scip.Signature{
					Language: "go",
					Text:     signature,
				}
			}
		}
		if hover := d.extractHoverText(parent, ident); hover != "" {
			documentation = append(documentation, hover)
		}
		if genDecl, ok := parent.(*ast.GenDecl); ok && genDecl.Doc != nil {
			blockDoc := strings.TrimSpace(genDecl.Doc.Text())
			if blockDoc != "" && (len(documentation) == 0 || documentation[0] != blockDoc) {
				documentation = append(documentation, blockDoc)
			}
		}
	}

	d.pkgSymbols.Set(pos, &scip.SymbolInformation{
		Symbol:                 symbol,
		Kind:                   symbols.KindForObject(def),
		DisplayName:            displayName,
		Documentation:          documentation,
		SignatureDocumentation: sigDoc,
		Relationships:          []*scip.Relationship{},
	})
}

func (d *Document) extractHoverText(parent ast.Node, node ast.Node) string {
	switch v := node.(type) {
	case *ast.File:
		if v.Doc != nil {
			return strings.TrimSpace(v.Doc.Text())
		} else {
			return fmt.Sprintf("package %s", v.Name.Name)
		}
	case *ast.FuncDecl:
		return strings.TrimSpace(v.Doc.Text())
	case *ast.GenDecl:
		return strings.TrimSpace(v.Doc.Text())
	case *ast.TypeSpec:
		// Typespecs do not have the doc associated with them much
		// of the time. They are often associated with the `type`
		// token itself.
		//
		// This is why we have to pass the declaration node
		doc := strings.TrimSpace(v.Doc.Text())
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.ValueSpec:
		doc := strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
		if doc == "" && parent != nil {
			doc = d.extractHoverText(nil, parent)
		}

		return doc
	case *ast.Field:
		return strings.TrimSpace(v.Doc.Text() + "\n" + v.Comment.Text())
	case *ast.Ident:
		if genDecl, ok := parent.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if s.Name == v {
						return d.extractHoverText(parent, s)
					}
				case *ast.ValueSpec:
					for _, name := range s.Names {
						if name == v {
							return d.extractHoverText(parent, s)
						}
					}
				}
			}
		}
		if parent != nil {
			return d.extractHoverText(nil, parent)
		}
	}

	return ""
}

// relativeQualifier returns a types.Qualifier that omits the package name for
// types in pkg (the "current" package) but includes it for all other packages.
// This ensures that e.g. embedded fields like io.Reader retain their package prefix.
func relativeQualifier(pkg *types.Package) types.Qualifier {
	return func(other *types.Package) string {
		if pkg == other {
			return ""
		}
		return other.Name()
	}
}

func typeStringForObject(obj types.Object) string {
	qual := relativeQualifier(obj.Pkg())

	switch v := obj.(type) {
	case *types.PkgName:
		return fmt.Sprintf("package %s", v.Name())

	case *types.TypeName:
		return formatTypeDeclaration(v)

	case *types.Var:
		if v.IsField() {
			// TODO(tjdevries) - make this be "(T).F" instead of "struct field F string"
			return fmt.Sprintf("struct %s", quotedTagsToBacktick(types.ObjectString(obj, qual)))
		}

	case *types.Const:
		return fmt.Sprintf("%s = %s", types.ObjectString(v, qual), v.Val())
	}

	return types.ObjectString(obj, qual)
}

var loggedGODEBUGWarning sync.Once

// formatTypeDeclaration returns the full type declaration string for a type,
// e.g. "type T struct{}", "type I interface { ... }", "type U = T", "type Z int32".
func formatTypeDeclaration(obj *types.TypeName) string {
	if obj.IsAlias() {
		return formatAliasDeclaration(obj)
	}

	return fmt.Sprintf("type %s %s", obj.Name(), expandTypeExpr(obj.Pkg(), obj.Type().Underlying()))
}

// formatAliasDeclaration returns the type declaration for alias types.
func formatAliasDeclaration(obj *types.TypeName) string {
	switch ty := obj.Type().(type) {
	case *types.Alias:
		qual := relativeQualifier(obj.Pkg())
		lhs := obj.Name() + formatTypeParamList(ty.TypeParams())
		rhs := ty.Rhs()
		switch rhs.(type) {
		case *types.Alias, *types.Named:
			return fmt.Sprintf("type %s = %s", lhs, types.TypeString(rhs, qual))
		default:
			return fmt.Sprintf("type %s = %s", lhs, expandTypeExpr(obj.Pkg(), rhs))
		}
	default:
		if val := os.Getenv("GODEBUG"); strings.Contains(val, "gotypealias=0") {
			loggedGODEBUGWarning.Do(func() {
				slog.Warn(
					"Running with GODEBUG=gotypealias=0, this may cause incorrect hover docs")
			})
		} else {
			slog.Warn(
				"IsAlias() is true but Type() is not Alias; please report this as a bug",
				"obj", obj.String(), "obj.Type()", ty.String())
		}
	}

	// Fallback for when GODEBUG=gotypealias=0 or unexpected types.
	return fmt.Sprintf("type %s %s", obj.Name(), expandTypeExpr(obj.Pkg(), obj.Type().Underlying()))
}

// formatTypeParamList renders a TypeParamList as "[K C1, V C2]".
// Returns "" when tparams is nil or empty.
func formatTypeParamList(tparams *types.TypeParamList) string {
	if tparams == nil || tparams.Len() == 0 {
		return ""
	}
	var buf strings.Builder
	buf.WriteByte('[')
	for i := range tparams.Len() {
		if i > 0 {
			buf.WriteString(", ")
		}
		tp := tparams.At(i)
		buf.WriteString(tp.Obj().Name())
		buf.WriteByte(' ')
		buf.WriteString(tp.Constraint().String())
	}
	buf.WriteByte(']')
	return buf.String()
}

// expandTypeExpr renders a type expression, formatting struct and interface
// types with aligned fields using go/format.
func expandTypeExpr(pkg *types.Package, t types.Type) string {
	raw := types.TypeString(t, relativeQualifier(pkg))

	switch t.(type) {
	case *types.Struct, *types.Interface:
		if formatted, err := formatGoDecl("type _ " + raw); err == nil {
			return strings.TrimPrefix(formatted, "type _ ")
		}
	}

	return raw
}

// formatGoDecl formats a Go type declaration using go/format, returning
// the formatted result with tabs replaced by spaces.
func formatGoDecl(decl string) (string, error) {
	src := "package p\n\n" + decl + "\n"
	formatted, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}
	result := strings.TrimPrefix(string(formatted), "package p\n\n")
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "\t", indent)
	return result, nil
}

// quotedTagsToBacktick replaces double-quoted struct tag strings with
// backtick-quoted equivalents in the output of types.ObjectString.
func quotedTagsToBacktick(s string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	for i := 0; i < len(s); i++ {
		if s[i] != '"' {
			buf.WriteByte(s[i])
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[j] == '\\' {
				j++
				continue
			}
			if s[j] == '"' {
				quoted := s[i : j+1]
				if unquoted, err := strconv.Unquote(quoted); err == nil {
					buf.WriteByte('`')
					buf.WriteString(unquoted)
					buf.WriteByte('`')
				} else {
					buf.WriteString(quoted)
				}
				i = j
				break
			}
		}
	}
	return buf.String()
}
