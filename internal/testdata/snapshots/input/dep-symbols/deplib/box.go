package deplib

type Box[T any] struct {
	Value T
}

func (b Box[T]) Get() T {
	return b.Value
}

type Config struct {
	Name    string
	Verbose bool
}

const DefaultName = "default"

var GlobalCounter int

type Stringer interface {
	String() string
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
