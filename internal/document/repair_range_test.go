package document

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/scip-code/scip/bindings/go/scip"
)

// docFor writes src to a temp file and returns a Document over it, the way
// NewDocument would.
func docFor(t *testing.T, src string) *Document {
	t.Helper()
	path := filepath.Join(t.TempDir(), "geometry.go")
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}
	return &Document{originAbs: path, lineLen: loadLineLengths(path)}
}

// The line cgo mangles: `C.puts` is rewritten to `_Cfunc_puts`, so the range is
// sized 11 instead of 6 and runs past the end of the real line.
const putsSrc = "package example\n\nfunc f(cs *C.char) {\n    C.puts(cs)\n}\n"

func TestRepairRangeRecoversMangledCgoName(t *testing.T) {
	d := docFor(t, putsSrc)
	// Line 3 is "    C.puts(cs)" (14 bytes); the mangled range ends at 15.
	oob := rng(3, 4, 3, 15)
	if d.InBounds(oob) {
		t.Fatal("test setup: range should be out of bounds")
	}

	got, ok := d.RepairRange(oob)
	if !ok {
		t.Fatal("RepairRange should recover a mangled cgo name")
	}
	if want := rng(3, 4, 3, 10); got != want {
		t.Errorf("RepairRange = %v, want %v (covering `C.puts`)", got, want)
	}
	if !d.InBounds(got) {
		t.Error("repaired range must be in bounds")
	}
}

func TestRepairRangeRejectsUnmeasurablePositions(t *testing.T) {
	d := docFor(t, putsSrc)

	cases := []struct {
		name string
		r    scip.Range
	}{
		// cgo's `defer C.f(x)` wrapper lands at end-of-line, where there is no
		// identifier to measure -- inventing a range there would be a guess.
		{"start at end of line", rng(3, 14, 3, 25)},
		{"start past end of line", rng(3, 40, 3, 51)},
		{"line past EOF", rng(99, 0, 99, 5)},
		{"multi-line range", rng(3, 4, 4, 2)},
		{"negative line", rng(-1, 0, -1, 5)},
		{"negative column", rng(3, -1, 3, 5)},
		// Column 0 of `    C.puts(cs)` is a space: no identifier starts there.
		{"start on non-identifier", rng(3, 0, 3, 40)},
	}
	for _, tc := range cases {
		if _, ok := d.RepairRange(tc.r); ok {
			t.Errorf("%s: RepairRange(%v) should not repair", tc.name, tc.r)
		}
	}
}

func TestRepairRangeUnreadableSource(t *testing.T) {
	// A document whose source can't be read has nothing to measure against.
	d := &Document{originAbs: filepath.Join(t.TempDir(), "missing.go")}
	if _, ok := d.RepairRange(rng(0, 0, 0, 5)); ok {
		t.Error("unreadable source should not repair")
	}
}

func TestRepairRangeHandlesNonASCIIIdentifiers(t *testing.T) {
	// Ranges are byte offsets, and Go identifiers may be non-ASCII, so the
	// repaired width must be measured in bytes.
	d := docFor(t, "package example\n\nvar été = 1\n")
	// Line 2 is "var été = 1"; `été` starts at byte 4 and is 5 bytes long.
	got, ok := d.RepairRange(rng(2, 4, 2, 99))
	if !ok {
		t.Fatal("should repair a non-ASCII identifier")
	}
	if want := rng(2, 4, 2, 9); got != want {
		t.Errorf("RepairRange = %v, want %v", got, want)
	}
}
