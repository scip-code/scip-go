package visitors

import (
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func TestSourceFileIdentifierAt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "source.go")
	if err := os.WriteFile(path, []byte("package p\n//line other.go:100\nvar x = 1\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	source := &sourceFile{path: path}

	if ident := source.identifierAt(token.Position{Filename: path, Line: 3, Column: 5}); ident == nil || ident.Name != "x" {
		t.Errorf("identifierAt(3:5) = %v, want x", ident)
	}
	// Positions are unadjusted, so the line directive has no effect.
	if ident := source.identifierAt(token.Position{Filename: path, Line: 100, Column: 5}); ident != nil {
		t.Errorf("identifierAt(100:5) = %v, want nil", ident)
	}

	if ident := source.identifierAt(token.Position{Filename: "other.go", Line: 3, Column: 5}); ident != nil {
		t.Errorf("identifierAt in another file = %v, want nil", ident)
	}

	missingPath := filepath.Join(t.TempDir(), "missing.go")
	missing := &sourceFile{path: missingPath}
	if ident := missing.identifierAt(token.Position{Filename: missingPath, Line: 1, Column: 1}); ident != nil {
		t.Errorf("identifierAt in a missing file = %v, want nil", ident)
	}
}
