# Bare `PkgName` identifier

`var _ = sort` is syntactically valid Go but ill-typed; the compiler rejects it
with `use of package sort not in selector`. The parser still accepts it and the
type checker, run in best-effort mode by `packages.Load`, still records
`info.Uses[sort] = *types.PkgName`.

Well-typed code routes such uses through the `*ast.SelectorExpr` handler in
`visitor_file.go`; a bare ident slips past it into the generic reference path
and previously hit the `panic("should never lookup PkgName ...")` assertion in
`lookup.go`. The visitor now skips bare `*types.PkgName` references instead of
crashing.
