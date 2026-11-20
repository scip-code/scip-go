package visitors

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestSourceMap(t *testing.T) {
	root := t.TempDir()
	goFile := filepath.Join(root, "source.go")
	fset := token.NewFileSet()
	source, err := parser.ParseFile(fset, goFile, "//line ignored.y:100\npackage example", 0)
	if err != nil {
		t.Fatal(err)
	}
	compiledOnly, err := parser.ParseFile(fset, filepath.Join(root, "generated.go"), "//line source.go:2:9\npackage example\n//line unknown.go:10:1\nvar x int", 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg := &packages.Package{Fset: fset, GoFiles: []string{goFile}}
	pkg.Syntax = append(pkg.Syntax, source, compiledOnly)
	sourceMap := NewSourceMap(pkg)

	for _, tc := range []struct {
		pos          token.Pos
		line, column int
	}{
		{source.Name.Pos(), 2, 9},
		{compiledOnly.Package, 2, 9},
	} {
		pos := sourceMap.Position(tc.pos)
		if pos.Filename != goFile || pos.Line != tc.line || pos.Column != tc.column {
			t.Errorf("Position = %v, want %s:%d:%d", pos, goFile, tc.line, tc.column)
		}
	}
	if pos := sourceMap.Position(compiledOnly.Decls[0].Pos()); pos.IsValid() {
		t.Errorf("Position outside pkg.GoFiles = %v, want zero value", pos)
	}
	if !sourceMap.InGoFiles(source) || sourceMap.InGoFiles(compiledOnly) {
		t.Error("InGoFiles should only hold for source.go")
	}
}
