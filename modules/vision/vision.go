package vision

import (
	"github.com/KennethGrace/gracious/base"
	"github.com/KennethGrace/gracious/modules"
)

type Vision struct {
	modules.Module
	feedback base.Feedback
}

func NewVision(vertical int, horizontal int) *Vision {
	return &Vision{}
}

func (v Vision) Begin() {
	//TODO implement me
	panic("implement me")
}
