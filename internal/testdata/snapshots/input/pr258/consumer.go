package consumer

import "github.com/example/deplib"

var Sentinel deplib.CustomErr

func New() deplib.CustomErr { return nil }

type Wrapper struct {
	Err deplib.CustomErr
}
