package visitors

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

// SourceMap maps positions in pkg.CompiledGoFiles back to pkg.GoFiles.
//
// For tools like cgo that generate Go files with //line directives,
// this gives us positions in the original source file ("adjusted").
type SourceMap struct {
	fset    *token.FileSet
	goFiles map[string]bool
}

func NewSourceMap(pkg *packages.Package) *SourceMap {
	s := &SourceMap{fset: pkg.Fset, goFiles: map[string]bool{}}
	for _, path := range pkg.GoFiles {
		s.goFiles[path] = true
	}
	return s
}

// Position returns the location of pos in the most relevant file.
//
// For normal code this is the source file itself. For generated code
// like CGo, we prefer the original foo.go position over the
// foo.cgo1.go position.
func (s *SourceMap) Position(pos token.Pos) token.Position {
	if unadjusted := s.fset.PositionFor(pos, false); s.goFiles[unadjusted.Filename] {
		return unadjusted
	}
	if adjusted := s.fset.PositionFor(pos, true); s.goFiles[adjusted.Filename] {
		return adjusted
	}
	return token.Position{}
}

// InGoFiles reports whether file was parsed from one of pkg.GoFiles.
func (s *SourceMap) InGoFiles(file *ast.File) bool {
	return s.goFiles[s.fset.File(file.Pos()).Name()]
}
