package q

import "sg/pr312/p"

type Impl struct{}

func (Impl) Run(o *p.Opts) error { return nil }

func (Impl) Nickname() string { return "" }
