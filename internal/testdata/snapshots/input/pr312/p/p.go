package p

type Opts struct{}

type Svc interface {
	Run(o *Opts) error
}
