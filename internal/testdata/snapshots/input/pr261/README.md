# Bare `PkgName` identifier

`var _ = sort` is malformed Go but the parser accepts it and the type checker
still records `info.Uses[sort] = *types.PkgName`. Well-formed code routes such
uses through the `*ast.SelectorExpr` handler in `visitor_file.go`; a bare ident
slips past it into the generic reference path and previously hit the
`panic("should never lookup PkgName ...")` assertion in `lookup.go`. The visitor
now skips bare `*types.PkgName` references instead of crashing.
