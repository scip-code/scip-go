package document

import (
	"testing"

	"github.com/scip-code/scip/bindings/go/scip"
)

func rng(sl, sc, el, ec int32) scip.Range {
	return scip.Range{
		Start: scip.Position{Line: sl, Character: sc},
		End:   scip.Position{Line: el, Character: ec},
	}
}

func TestInBounds(t *testing.T) {
	// Two lines, byte lengths 5 and 10.
	d := &Document{lineLen: []int{5, 10}}

	cases := []struct {
		name string
		r    scip.Range
		want bool
	}{
		{"in bounds, single line", rng(0, 0, 0, 5), true},
		{"end column at EOL", rng(1, 0, 1, 10), true},
		{"spans two lines", rng(0, 1, 1, 2), true},
		{"start column past EOL", rng(0, 6, 0, 6), false},
		{"end column past EOL", rng(1, 0, 1, 11), false},
		{"line past EOF", rng(2, 0, 2, 0), false},
		{"negative start column", rng(0, -1, 0, -1), false},
		{"negative line", rng(-1, 0, -1, 0), false},
		{"reversed same line", rng(0, 4, 0, 1), false},
		{"reversed across lines", rng(1, 0, 0, 0), false},
	}
	for _, tc := range cases {
		if got := d.InBounds(tc.r); got != tc.want {
			t.Errorf("%s: InBounds(%v) = %v, want %v", tc.name, tc.r, got, tc.want)
		}
	}
}

func TestInBoundsUnknownSourceAdmitsWellFormed(t *testing.T) {
	// A document whose source couldn't be read (lineLen == nil) must not
	// over-drop: it admits any well-formed range but still rejects malformed ones.
	d := &Document{}
	if !d.InBounds(rng(999, 0, 999, 3)) {
		t.Error("nil lineLen should admit a well-formed range")
	}
	if d.InBounds(rng(0, -1, 0, 0)) {
		t.Error("nil lineLen should still reject a negative column")
	}
	if d.InBounds(rng(2, 0, 1, 0)) {
		t.Error("nil lineLen should still reject a reversed range")
	}
}
