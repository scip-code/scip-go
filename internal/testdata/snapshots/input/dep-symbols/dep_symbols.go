package depsymbols

import "github.com/example/deplib"

func UseGenericField() int {
	b := deplib.Box[int]{Value: 42}
	return b.Value
}

func UseGenericMethod() string {
	b := deplib.Box[string]{Value: "hello"}
	return b.Get()
}

func UseNonGenericField() string {
	c := deplib.Config{Name: "test", Verbose: true}
	return c.Name
}

func UseConst() string {
	return deplib.DefaultName
}

func UseVar() int {
	return deplib.GlobalCounter
}

type LocalType struct{}

func (l LocalType) String() string { return "local" }

type EmbeddedStringer struct {
	LocalType
}

type LocalInterface interface {
	Get() int
}

func UseDepWriter(w deplib.Writer) {
	w.Write(nil)
}
