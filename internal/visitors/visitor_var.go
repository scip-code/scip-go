package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"log/slog"

	"github.com/scip-code/scip-go/internal/document"
	"github.com/scip-code/scip-go/internal/symbols"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitVarDefinition(doc *document.Document, pkg *packages.Package, decl *ast.GenDecl) {
	ast.Walk(&varVisitor{
		doc:   doc,
		pkg:   pkg,
		scope: NewScope(pkg.PkgPath),
	}, decl)
}

type varVisitor struct {
	doc *document.Document
	pkg *packages.Package

	curDecl ast.Decl
	scope   *Scope
}

var _ ast.Visitor = &varVisitor{}

func (v *varVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.GenDecl:
		switch node.Tok {
		// Only traverse vars and consts
		case token.VAR, token.CONST:
			v.curDecl = node
			return v
		default:
			return nil
		}

	case *ast.ValueSpec:
		// Iterate over names, which are the only thing that can be definitions
		for _, name := range node.Names {
			symbol := v.makeSymbol(descriptorTerm(name.Name))
			v.doc.SetNewSymbol(symbol, v.curDecl, name)
		}

		if len(node.Names) > 0 {
			var scopeName string
			if len(node.Names) == 1 {
				scopeName = node.Names[0].Name
			} else {
				position := v.pkg.Fset.Position(node.Pos())
				scopeName = fmt.Sprintf("inline-%d-%d", position.Line, position.Column)
			}

			v.scope.push(scopeName, scip.Descriptor_Meta)
			if node.Type != nil {
				switch valueType := node.Type.(type) {
				case *ast.FuncType:
					// nothing
				case *ast.Ident:
					// nothing
				case *ast.ChanType:
					// TODO: Could have chan with new struct
				case *ast.ArrayType:
					// TODO: Could have array with new struct
				case *ast.MapType:
					// TODO: could have nested struct?
				case *ast.StarExpr:
					// TODO: could be *struct?
				case *ast.SelectorExpr:
					// TODO?
				case *ast.IndexExpr:
					// TODO: Generics, possibly need to traverse
				case *ast.IndexListExpr:
					// TODO: Generics, possibly need to traverse
				case *ast.InterfaceType:
					ast.Walk(v, valueType)
				case *ast.StructType:
					// panic(fmt.Sprintf("TODO: handle type %T %s", valueType, v.pkg.Fset.Position(node.Pos())))
					ast.Walk(v, valueType)
				default:
					slog.Debug(fmt.Sprintf("TODO: handle type %T %s", valueType, v.pkg.Fset.Position(node.Pos())))

					return
				}
			}

			walkExprList(v, node.Values)
			v.scope.pop()
		}

		return nil

	case *ast.Field:
		// I think the only case of this is embedded fields.
		if len(node.Names) == 0 {

			names := getIdentOfTypeExpr(v.pkg.Fset, node.Type)
			for _, name := range names {
				embeddedSymbol := v.makeSymbol(&scip.Descriptor{
					Name:   name.Name,
					Suffix: scip.Descriptor_Term,
				})

				// In this odd scenario, the definition is at embedded field level,
				// not wherever the name is. So that messes up our lookup table.
				v.doc.SetNewSymbolForPos(embeddedSymbol, node, name, node.Pos())
			}
		} else {
			for _, name := range node.Names {
				v.doc.SetNewSymbol(v.makeSymbol(&scip.Descriptor{
					Name:   name.Name,
					Suffix: scip.Descriptor_Term,
				}), node, name)

				v.walkAnonymousTypeFields(node.Type)
			}
		}
		return nil

	default:
		return v
	}
}

func (v *varVisitor) walkAnonymousTypeFields(expr ast.Expr) {
	switch typ := expr.(type) {
	case *ast.MapType:
		v.walkAnonymousTypeFields(typ.Key)
		v.walkAnonymousTypeFields(typ.Value)
	case *ast.ArrayType:
		v.walkAnonymousTypeFields(typ.Elt)
	case *ast.StarExpr:
		v.walkAnonymousTypeFields(typ.X)
	case *ast.StructType, *ast.InterfaceType:
		if key := anonTypeKey(v.pkg, expr); key != "" {
			v.scope.withScope(key, scip.Descriptor_Type, func() {
				ast.Walk(v, expr)
			})
		}
	}
}

func (s *varVisitor) makeSymbol(descriptor *scip.Descriptor) string {
	return symbols.FromDescriptors(s.pkg, append(s.scope.descriptors, descriptor)...)
}
