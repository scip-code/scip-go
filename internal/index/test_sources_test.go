package index_test

import (
	"os"
	"path/filepath"
	"slices"
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
