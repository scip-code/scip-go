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

func VisitPackageSyntax(
	moduleRoot string,
	pkg *packages.Package,
	pathToDocuments map[string]*document.Document,
	globalSymbols *lookup.Global,
) {
	pkgSymbols := lookup.NewPackageSymbols(pkg)
	// Iterate over all the files, collect any global symbols
	for _, f := range pkg.Syntax {

		abs := pkg.Fset.File(f.Package).Name()
		relative, _ := filepath.Rel(moduleRoot, abs)

		doc := visitSyntax(pkg, pkgSymbols, f, relative)

		// Save document for pass 2
		pathToDocuments[abs] = doc
	}

	globalSymbols.Add(pkgSymbols)
}

func visitSyntax(pkg *packages.Package, pkgSymbols *lookup.Package, f *ast.File, relative string) *document.Document {
	doc := document.NewDocument(relative, pkg, pkgSymbols)

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

func scipRange(start, end token.Position, obj types.Object) []int32 {
	var adjustment int32 = 0
	if pkgName, ok := obj.(*types.PkgName); ok && strings.HasPrefix(pkgName.Name(), `"`) {
		adjustment = 1
	}

	startLine := int32(start.Line - 1)
	startColumn := int32(start.Column - 1)
	endLine := int32(end.Line - 1)
	endColumn := int32(end.Column - 1)
	if startLine != endLine {
		return []int32{startLine, startColumn + adjustment, endLine, endColumn - adjustment}
	}
	return []int32{startLine, startColumn + adjustment, endColumn - adjustment}
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
