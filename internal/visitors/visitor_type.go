package visitors

import (
	"go/ast"

	"github.com/scip-code/scip-go/internal/document"
	"github.com/scip-code/scip-go/internal/symbols"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func visitTypesInFile(doc *document.Document, pkg *packages.Package, file *ast.File) {
	visitor := typeVisitor{
		pkg:   pkg,
		doc:   doc,
		scope: NewScope(pkg.PkgPath),
	}

	ast.Walk(visitor, file)
}

// typeVisitor collects the all the information for top-level structs
// that can be imported by any other file (they do not have to be exported).
//
// For example, a struct `myStruct` can be imported by other files in the same
// packages. So we need to make those field names global (we only have global
// or file-local).
type typeVisitor struct {
	doc *document.Document
	pkg *packages.Package

	scope       *Scope
	curDecl     *ast.GenDecl
	isInterface bool
}

func (v typeVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.GenDecl:
		// Current declaration is required for some documentation parsing.
		// So we have to keep this here with us as we traverse more deeply
		v.curDecl = node

		if node.Doc != nil {
			ast.Walk(v, node.Doc)
		}

		for _, s := range node.Specs {
			switch spec := s.(type) {
			case *ast.TypeSpec:
				switch spec.Type.(type) {
				case *ast.InterfaceType:
					v.isInterface = true
				default:
					v.isInterface = false
				}
			}

			ast.Walk(v, s)
		}

		return nil

	case
		// Continue down file and decls
		*ast.File,

		// Toplevel types that are important
		*ast.StructType,
		*ast.InterfaceType,

		// Continue traversing subtypes
		*ast.FieldList,
		*ast.Ident:

		return v

	case *ast.TypeSpec:
		v.scope.push(node.Name.Name, scip.Descriptor_Type)
		defer func() {
			v.scope.pop()
		}()

		v.doc.SetNewSymbol(
			symbols.FromDescriptors(v.pkg, v.scope.descriptors...),
			v.curDecl,
			node.Name,
		)

		ast.Walk(v, node.Type)
	case *ast.Field:
		// I think the only case of this is embedded fields.
		if len(node.Names) == 0 {
			// If we have an interface, these do not *declare* a new symbol,
			// they simply add another constraint.
			if v.isInterface {
				return v
			}

			names := getIdentOfTypeExpr(v.pkg.Fset, node.Type)
			for _, name := range names {
				embeddedSymbol := v.makeSymbol(&scip.Descriptor{
					Name:   name.Name,
					Suffix: scip.Descriptor_Term,
				})

				// In this odd scenario, the definition is at embedded field level,
				// not wherever the name is. So that messes up our lookup table.
				v.doc.SetNewSymbolForPos(embeddedSymbol, node, name, name.Pos())
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
	}

	return nil
}

// Implements ast.Visitor
var _ ast.Visitor = &typeVisitor{}

func (v typeVisitor) walkAnonymousTypeFields(expr ast.Expr) {
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

func (s *typeVisitor) makeSymbol(descriptor *scip.Descriptor) string {
	return symbols.FromDescriptors(s.pkg, append(s.scope.descriptors, descriptor)...)
}
