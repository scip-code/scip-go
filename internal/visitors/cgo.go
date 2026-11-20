package visitors

import (
	"go/ast"
	"slices"

	"github.com/scip-code/scip-go/internal/lookup"
	"github.com/scip-code/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

// cgoPrefixes are the prefixes cmd/cgo adds to the Go names it generates for C.name.
var cgoPrefixes = []string{"_Cfunc_", "_Ctype_", "_Cvar_", "_Ciconst_", "_Cfconst_", "_Csconst_", "_Cmacro_"}

// cgoReferences returns references from C.name selectors in the file to
// the Go declarations that cmd/cgo generates for them. The compiled-only
// files cannot provide these, because cmd/cgo rewrites the selectors.
func (s *sourceFile) cgoReferences(pkg *packages.Package, global *lookup.Global) []*scip.Occurrence {
	file := s.parse()
	if file == nil || !slices.ContainsFunc(file.Imports, func(spec *ast.ImportSpec) bool {
		return spec.Path.Value == `"C"`
	}) {
		return nil
	}

	// cmd/cgo uses the errno wrapper when a single call supplies two values.
	errnoCalls := map[*ast.SelectorExpr]bool{}
	markErrnoCall := func(lhsCount int, rhs []ast.Expr) {
		if lhsCount == 2 && len(rhs) == 1 {
			if call, ok := ast.Unparen(rhs[0]).(*ast.CallExpr); ok {
				if selector, ok := ast.Unparen(call.Fun).(*ast.SelectorExpr); ok {
					errnoCalls[selector] = true
				}
			}
		}
	}

	var occurrences []*scip.Occurrence
	ast.Inspect(file, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.AssignStmt:
			markErrnoCall(len(n.Lhs), n.Rhs)
		case *ast.ValueSpec:
			markErrnoCall(len(n.Names), n.Values)
		}
		selector, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		// A resolved C refers to a local declaration, not the cgo pseudo-package.
		if x, ok := selector.X.(*ast.Ident); !ok || x.Name != "C" || x.Obj != nil {
			return true
		}
		name := selector.Sel.Name
		if name == "malloc" {
			// cmd/cgo renames C.malloc to C._CMalloc in ast.go.
			name = "_CMalloc"
		}
		prefixes := cgoPrefixes
		if errnoCalls[selector] {
			prefixes = []string{"_C2func_"}
		}
		for _, prefix := range prefixes {
			obj := pkg.Types.Scope().Lookup(prefix + name)
			if obj == nil {
				continue
			}
			if info, ok, err := global.GetSymbolOfObject(obj); err == nil && ok {
				occurrences = append(occurrences, &scip.Occurrence{
					Symbol:      info.Symbol,
					SymbolRoles: int32(scip.SymbolRole_ReadAccess),
					TypedRange: scipRange(s.fset.PositionFor(selector.Pos(), false),
						s.fset.PositionFor(selector.End(), false), nil).AsTypedRange(),
				})
				break
			}
		}
		return true
	})
	return occurrences
}
