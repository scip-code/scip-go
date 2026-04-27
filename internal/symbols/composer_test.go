package symbols_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/scip-code/scip-go/internal/symbols"
	"golang.org/x/tools/go/packages"
)

func TestComposeNil(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	if got := c.Compose(pkg, nil); got != "" {
		t.Errorf("Compose(nil) = %q, want empty", got)
	}
}

func TestComposeNilPkg(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	obj := types.Universe.Lookup("len")
	if got := c.Compose(pkg, obj); got != "" {
		t.Errorf("Compose(builtin) = %q, want empty", got)
	}
}

func TestComposeTypeName(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	obj := types.NewTypeName(token.NoPos, tpkg, "MyStruct", nil)
	types.NewNamed(obj, types.NewStruct(nil, nil), nil)

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/MyStruct#"
	if got != want {
		t.Errorf("Compose(TypeName) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeConst(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	obj := types.NewConst(token.NoPos, tpkg, "MaxRetries", types.Typ[types.Int], nil)

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/MaxRetries."
	if got != want {
		t.Errorf("Compose(Const) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposePackageVar(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	obj := types.NewVar(token.NoPos, tpkg, "GlobalCount", types.Typ[types.Int])

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/GlobalCount."
	if got != want {
		t.Errorf("Compose(Var) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeFunc(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	sig := types.NewSignatureType(nil, nil, nil, nil, nil, false)
	obj := types.NewFunc(token.NoPos, tpkg, "DoWork", sig)

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/DoWork()."
	if got != want {
		t.Errorf("Compose(Func) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeMethod(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	typeName := types.NewTypeName(token.NoPos, tpkg, "Server", nil)
	named := types.NewNamed(typeName, types.NewStruct(nil, nil), nil)

	recv := types.NewVar(token.NoPos, tpkg, "s", named)
	sig := types.NewSignatureType(recv, nil, nil, nil, nil, false)
	obj := types.NewFunc(token.NoPos, tpkg, "Start", sig)

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/Server#Start()."
	if got != want {
		t.Errorf("Compose(Method) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeMethodPointerReceiver(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	typeName := types.NewTypeName(token.NoPos, tpkg, "Server", nil)
	named := types.NewNamed(typeName, types.NewStruct(nil, nil), nil)

	recv := types.NewVar(token.NoPos, tpkg, "s", types.NewPointer(named))
	sig := types.NewSignatureType(recv, nil, nil, nil, nil, false)
	obj := types.NewFunc(token.NoPos, tpkg, "Stop", sig)

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/Server#Stop()."
	if got != want {
		t.Errorf("Compose(Method ptr recv) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeStructField(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	field := types.NewField(token.NoPos, tpkg, "Name", types.Typ[types.String], false)
	structType := types.NewStruct([]*types.Var{field}, nil)

	typeName := types.NewTypeName(token.NoPos, tpkg, "Config", nil)
	types.NewNamed(typeName, structType, nil)
	tpkg.Scope().Insert(typeName)

	got := c.Compose(pkg, field)
	want := "scip-go gomod example.com/lib v1.0.0 `example.com/lib`/Config#Name."
	if got != want {
		t.Errorf("Compose(Field) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeDefaultModule(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{
		DefaultModulePath:    "example.com/project",
		DefaultModuleVersion: "1.0.0",
	})
	pkg := &packages.Package{PkgPath: "example.com/lib"}

	tpkg := types.NewPackage("example.com/lib", "lib")
	obj := types.NewConst(token.NoPos, tpkg, "Version", types.Typ[types.String], nil)

	got := c.Compose(pkg, obj)
	want := "scip-go gomod example.com/project 1.0.0 `example.com/lib`/Version."
	if got != want {
		t.Errorf("Compose(default module) =\n  %q\nwant\n  %q", got, want)
	}
}

func TestComposeMemoization(t *testing.T) {
	c := symbols.NewComposer(symbols.ComposerConfig{})
	pkg := &packages.Package{PkgPath: "example.com/lib", Module: &packages.Module{Path: "example.com/lib", Version: "v1.0.0"}}

	tpkg := types.NewPackage("example.com/lib", "lib")
	obj := types.NewConst(token.NoPos, tpkg, "X", types.Typ[types.Int], nil)

	first := c.Compose(pkg, obj)
	second := c.Compose(pkg, obj)
	if first != second {
		t.Errorf("memoization broken: %q != %q", first, second)
	}
}
