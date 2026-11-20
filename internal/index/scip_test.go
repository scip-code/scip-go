package index_test

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scip-code/scip-go/internal/config"
	"github.com/scip-code/scip-go/internal/index"
	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/scip-code/scip/bindings/go/scip/testutil"
	"golang.org/x/mod/modfile"
	"google.golang.org/protobuf/proto"
)

// Use "update-snapshots" to update snapshots
var filter = flag.String("filter", "", "filenames to filter by")

func TestSnapshots(t *testing.T) {
	snapshotRoot := getTestdataRoot(t)

	testutil.SnapshotTest(t,
		snapshotRoot,
		func(inputDirectory, outputDirectory string, sources []*scip.SourceFile) []*scip.SourceFile {
			fmt.Println("Indexing", inputDirectory, "outputDirectory", outputDirectory)

			if *filter != "" && !strings.Contains(inputDirectory, *filter) {
				return []*scip.SourceFile{}
			}

			scipIndex := scip.Index{}
			writer := func(msg proto.Message) error {
				switch msg := msg.(type) {
				case *scip.Metadata:
					scipIndex.Metadata = msg
				case *scip.Document:
					scipIndex.Documents = append(scipIndex.Documents, msg)
				case *scip.SymbolInformation:
					scipIndex.ExternalSymbols = append(scipIndex.ExternalSymbols, msg)
				}
				return nil
			}

			opts := config.IndexOpts{
				ModuleRoot:      inputDirectory,
				ModuleVersion:   "0.1.test",
				ModulePath:      "sg/" + filepath.Base(inputDirectory),
				GoStdlibVersion: "go1.22",
			}

			driverDir := filepath.Join(inputDirectory, "driver")
			var driverPath string
			if _, err := os.Stat(filepath.Join(driverDir, "main.go")); err == nil {
				driverPath = buildGoPackagesDriver(t, driverDir)
				os.Setenv("GOPACKAGESDRIVER", driverPath)
				defer os.Unsetenv("GOPACKAGESDRIVER")
				opts.IsGoPackagesDriverSet = true
			}

			if err := index.Index(writer, opts); err != nil {
				t.Fatal(err)
			}

			if driverPath != "" {
				if _, err := os.Stat(driverPath + ".sentinel"); err != nil {
					t.Fatal("GOPACKAGESDRIVER was not invoked")
				}
			}

			// Index nested modules (such as dependencies wired in with a local
			// replace directive) on their own, as their published index would be,
			// and check that references to them match their definitions.
			for _, dep := range indexNestedModules(t, inputDirectory) {
				for _, doc := range scipIndex.Documents {
					for _, occ := range doc.Occurrences {
						if scip.SymbolRole_Definition.Matches(occ) || dep.definitions[occ.Symbol] {
							continue
						}
						// A replacement module with a different module path keeps the
						// replaced import paths, which its own index cannot match.
						sym, err := scip.ParseSymbol(occ.Symbol)
						if err == nil && sym.Package.GetName() == dep.modulePath && strings.HasPrefix(sym.Descriptors[0].Name, dep.modulePath) {
							t.Errorf("%s: reference %q has no definition in %s", doc.RelativePath, occ.Symbol, dep.modulePath)
						}
					}
				}
				scipIndex.Documents = append(scipIndex.Documents, dep.documents...)
			}

			// Filter out documents outside of current directory
			var filteredDocs []*scip.Document
			for _, doc := range scipIndex.Documents {
				if strings.HasPrefix(doc.RelativePath, "..") {
					continue
				}
				if *filter != "" && !strings.Contains(doc.RelativePath, *filter) {
					continue
				}
				filteredDocs = append(filteredDocs, doc)
			}
			scipIndex.Documents = filteredDocs
			assertTypedOccurrenceRanges(t, &scipIndex)

			symbolFormatter := scip.SymbolFormatter{
				OnError:               func(err error) error { return err },
				IncludeScheme:         func(scheme string) bool { return scheme == "local" },
				IncludePackageManager: func(_ string) bool { return false },
				IncludePackageName:    func(name string) bool { return !strings.HasPrefix(name, "sg/") },
				IncludePackageVersion: func(_ string) bool { return true },
				IncludeDescriptor:     func(_ string) bool { return true },
				IncludeRawDescriptor:  func(descriptor *scip.Descriptor) bool { return true },
				IncludeDisambiguator:  func(_ string) bool { return true },
			}

			sourceFiles, err := testutil.FormatSnapshots(
				&scipIndex, "//", symbolFormatter, "")
			if err != nil {
				t.Fatal(err)
			}

			return sourceFiles
		},
	)
}

type nestedModule struct {
	modulePath  string
	documents   []*scip.Document
	definitions map[string]bool
}

// indexNestedModules indexes every module below root, except the
// GOPACKAGESDRIVER in driver/, with document paths relative to root.
func indexNestedModules(t *testing.T, root string) []nestedModule {
	t.Helper()
	var modules []nestedModule
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.Name() != "go.mod" || filepath.Dir(path) == root {
			return err
		}
		dir := filepath.Dir(path)
		rel, err := filepath.Rel(root, dir)
		if err != nil || rel == "driver" {
			return err
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		module := nestedModule{modulePath: modfile.ModulePath(contents), definitions: map[string]bool{}}
		err = index.Index(func(msg proto.Message) error {
			if doc, ok := msg.(*scip.Document); ok {
				doc.RelativePath = filepath.Join(rel, doc.RelativePath)
				for _, occ := range doc.Occurrences {
					if scip.SymbolRole_Definition.Matches(occ) {
						module.definitions[occ.Symbol] = true
					}
				}
				module.documents = append(module.documents, doc)
			}
			return nil
		}, config.IndexOpts{
			ModuleRoot:      dir,
			ModuleVersion:   "0.1.test",
			ModulePath:      module.modulePath,
			GoStdlibVersion: "go1.22",
		})
		modules = append(modules, module)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return modules
}

func assertTypedOccurrenceRanges(t *testing.T, scipIndex *scip.Index) {
	t.Helper()
	for _, doc := range scipIndex.Documents {
		if doc.PositionEncoding != scip.PositionEncoding_UTF8CodeUnitOffsetFromLineStart {
			t.Fatalf("%s: position encoding = %v; want UTF-8 bytes", doc.RelativePath, doc.PositionEncoding)
		}
		for i, occ := range doc.Occurrences {
			if occ == nil {
				t.Fatalf("%s: nil occurrence at index %d", doc.RelativePath, i)
			}
			if occ.GetTypedRange() == nil {
				t.Fatalf("%s: occurrence %q has no typed range", doc.RelativePath, occ.GetSymbol())
			}
			if _, ok := occ.SourceRange(); !ok {
				t.Fatalf("%s: occurrence %q has malformed typed range", doc.RelativePath, occ.GetSymbol())
			}
			if len(occ.GetRange()) != 0 {
				t.Fatalf("%s: occurrence %q set deprecated range: %v", doc.RelativePath, occ.GetSymbol(), occ.GetRange())
			}
			if len(occ.GetEnclosingRange()) != 0 {
				t.Fatalf("%s: occurrence %q set deprecated enclosing_range: %v", doc.RelativePath, occ.GetSymbol(), occ.GetEnclosingRange())
			}
		}
	}
}

// getTestdataRoot returns the absolute path to the testdata directory of this repository.
func getTestdataRoot(t *testing.T) string {
	testdata, err := filepath.Abs("../testdata")
	if err != nil {
		t.Fatalf("unexpected error getting absolute directory: %s", err)
	}

	return testdata
}

// buildGoPackagesDriver compiles the GOPACKAGESDRIVER binary found in
// driverDir (a directory containing main.go and go.mod) and returns the
// path to the built binary.
func buildGoPackagesDriver(t *testing.T, driverDir string) string {
	t.Helper()

	driverPath := filepath.Join(t.TempDir(), "gopackagesdriver")

	cmd := exec.Command("go", "build", "-o", driverPath, ".")
	cmd.Dir = driverDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("building GOPACKAGESDRIVER in %s: %v\n%s", driverDir, err, out)
	}

	return driverPath
}
