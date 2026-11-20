package visitors

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/scip-code/scip/bindings/go/scip"
)

// sourceFile is a file in pkg.GoFiles that isn't in
// pkg.CompiledGoFiles.
//
// It's a Go file that's an input to build-time preprocessor like CGo.
type sourceFile struct {
	path         string
	file         *ast.File
	fset         *token.FileSet
	identifiers  map[lineColumn]*ast.Ident
	declarations map[token.Pos]*scip.Range
}

type lineColumn struct{ line, column int }

func (s *sourceFile) parse() *ast.File {
	if s.fset == nil {
		s.fset = token.NewFileSet()
		s.file, _ = parser.ParseFile(s.fset, s.path, nil, parser.ParseComments)
		s.identifiers = map[lineColumn]*ast.Ident{}
		if s.file != nil {
			ast.Inspect(s.file, func(node ast.Node) bool {
				if ident, ok := node.(*ast.Ident); ok {
					pos := s.fset.PositionFor(ident.Pos(), false)
					s.identifiers[lineColumn{pos.Line, pos.Column}] = ident
				}
				return true
			})
		}
	}
	return s.file
}

func (s *sourceFile) identifierAt(pos token.Position) *ast.Ident {
	if pos.Filename != s.path {
		return nil
	}
	s.parse()
	return s.identifiers[lineColumn{pos.Line, pos.Column}]
}

// hasIdentifier reports whether name is written at pos.
func (s *sourceFile) hasIdentifier(pos token.Position, name string) bool {
	ident := s.identifierAt(pos)
	return ident != nil && ident.Name == name
}

// hasImport reports whether the import path literal starts at pos.
func (s *sourceFile) hasImport(pos token.Position, literal string) bool {
	file := s.parse()
	if file == nil || pos.Filename != s.path {
		return false
	}
	for _, spec := range file.Imports {
		start := s.fset.PositionFor(spec.Path.Pos(), false)
		if spec.Path.Value == literal && start.Line == pos.Line && start.Column == pos.Column {
			return true
		}
	}
	return false
}

func (s *sourceFile) enclosingRange(pos token.Position) *scip.Range {
	ident := s.identifierAt(pos)
	if ident == nil {
		return nil
	}
	if s.declarations == nil {
		s.declarations = collectDeclarationRanges(s.fset, s.file)
	}
	return s.declarations[ident.Pos()]
}
