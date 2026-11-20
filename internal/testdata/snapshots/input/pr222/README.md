# On-demand dependency symbol resolution

Tests that references into a dependency module produce correct symbols when the
dependency's symbols are resolved on demand by the `Composer` rather than
eagerly indexed.

Exercises the major symbol shapes a dependency can expose: generic and
non-generic struct fields, methods on generic types, package-level consts and
vars, embedded fields, and interface methods used through a dependency-defined
interface, including generic interfaces and methods of embedded interfaces.

The `deplib/` subdirectory is a separate module wired in via a `replace`
directive in `go.mod`.
