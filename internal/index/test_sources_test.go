package index_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/scip-code/scip-go/internal/config"
	"github.com/scip-code/scip-go/internal/index"
	"github.com/scip-code/scip/bindings/go/scip"
	"google.golang.org/protobuf/proto"
)

func TestIndexTestPackages(t *testing.T) {
	root := t.TempDir()
	files := map[string]string{
		"go.mod":           "module example.test\ngo 1.26\n",
		"main.go":          "package main\nfunc main() {}\nfunc F() {}\n",
		"main_test.go":     "package main\nimport \"testing\"\nfunc TestF(t *testing.T) { F() }\n",
		"external_test.go": "package main_test\nimport \"testing\"\nfunc TestExternal(t *testing.T) {}\n",
	}
	for name, source := range files {
		if err := os.WriteFile(filepath.Join(root, name), []byte(source), 0600); err != nil {
			t.Fatal(err)
		}
	}
	for _, skipTests := range []bool{false, true} {
		t.Run(map[bool]string{false: "with tests", true: "without tests"}[skipTests], func(t *testing.T) {
			var paths []string
			err := index.Index(func(msg proto.Message) error {
				if doc, ok := msg.(*scip.Document); ok {
					paths = append(paths, doc.RelativePath)
				}
				return nil
			}, config.IndexOpts{ModuleRoot: root, ModuleVersion: "test", SkipTests: skipTests, SkipImplementations: true})
			if err != nil {
				t.Fatal(err)
			}
			slices.Sort(paths)
			want := []string{"external_test.go", "main.go", "main_test.go"}
			if skipTests {
				want = []string{"main.go"}
			}
			if !slices.Equal(paths, want) {
				t.Fatalf("documents = %v, want %v", paths, want)
			}
		})
	}
}

// TestTestRole checks that occurrences in in-package and external test files
// have the Test role, and that occurrences in non-test files do not.
func TestTestRole(t *testing.T) {
	root := t.TempDir()
	files := map[string]string{
		"go.mod":           "module example.test/lib\ngo 1.26\n",
		"lib.go":           "package lib\nfunc Add(a, b int) int { return a + b }\nfunc Double(a int) int { return Add(a, a) }\n",
		"lib_test.go":      "package lib\nimport \"testing\"\nfunc helper() int { return Add(1, 2) }\nfunc TestAdd(t *testing.T) { helper() }\n",
		"external_test.go": "package lib_test\nimport (\n\t\"testing\"\n\t\"example.test/lib\"\n)\nfunc TestExternal(t *testing.T) { lib.Add(1, 2) }\n",
	}
	for name, source := range files {
		if err := os.WriteFile(filepath.Join(root, name), []byte(source), 0600); err != nil {
			t.Fatal(err)
		}
	}
	var scipIndex scip.Index
	err := index.Index(func(msg proto.Message) error {
		if doc, ok := msg.(*scip.Document); ok {
			scipIndex.Documents = append(scipIndex.Documents, doc)
		}
		return nil
	}, config.IndexOpts{ModuleRoot: root, ModuleVersion: "test", SkipImplementations: true})
	if err != nil {
		t.Fatal(err)
	}
	assertTestRoles(t, &scipIndex)

	// Each case must match exactly one occurrence, so this also checks that
	// loading test variants of the package does not duplicate occurrences.
	for _, want := range []struct {
		path   string
		symbol string // suffix of the occurrence's symbol
		roles  scip.SymbolRole
	}{
		// Non-test code: the definition of Add and a reference to it.
		{"lib.go", "/Add().", scip.SymbolRole_Definition},
		{"lib.go", "/Add().", scip.SymbolRole_ReadAccess},
		// A test function, a helper and a reference to non-test code, in the
		// package itself and in an external test package.
		{"lib_test.go", "/TestAdd().", scip.SymbolRole_Definition | scip.SymbolRole_Test},
		{"lib_test.go", "/helper().", scip.SymbolRole_Definition | scip.SymbolRole_Test},
		{"lib_test.go", "/helper().", scip.SymbolRole_ReadAccess | scip.SymbolRole_Test},
		{"lib_test.go", "/Add().", scip.SymbolRole_ReadAccess | scip.SymbolRole_Test},
		{"external_test.go", "/TestExternal().", scip.SymbolRole_Definition | scip.SymbolRole_Test},
		{"external_test.go", "/Add().", scip.SymbolRole_ReadAccess | scip.SymbolRole_Test},
	} {
		wantDefinition := want.roles&scip.SymbolRole_Definition != 0
		var got []scip.SymbolRole
		for _, doc := range scipIndex.Documents {
			if doc.RelativePath != want.path {
				continue
			}
			for _, occ := range doc.Occurrences {
				if strings.HasSuffix(occ.Symbol, want.symbol) &&
					scip.SymbolRole_Definition.Matches(occ) == wantDefinition {
					got = append(got, scip.SymbolRole(occ.SymbolRoles))
				}
			}
		}
		if len(got) != 1 || got[0] != want.roles {
			t.Errorf("%s %s: symbol_roles = %d, want [%d]", want.path, want.symbol, got, want.roles)
		}
	}
}
