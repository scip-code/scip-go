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

func assertTypedOccurrenceRanges(t *testing.T, scipIndex *scip.Index) {
	t.Helper()
	for _, doc := range scipIndex.Documents {
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
