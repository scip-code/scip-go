# Unparseable syntax files

`packages.Load` returns a file with no `package` declaration (`no_package.go`)
as a zero-decl `*ast.File` whose `f.Package == NoPos`, which is not registered
in `pkg.Fset`. Previously `pkg.Fset.File(f.Package).Name()` panicked on the
resulting nil `*token.File`. The loader now drops such entries from
`pkg.Syntax`, so the valid sibling (`pr260.go`) is indexed normally.
