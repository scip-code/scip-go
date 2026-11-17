package p

type Named interface {
	Nickname() string
}

// NamedPtr is a constraint: only *T for some T with a Nickname method satisfies it.
// No named type implements it, so it must not get implementation edges.
type NamedPtr[T any] interface {
	Named
	*T
}
