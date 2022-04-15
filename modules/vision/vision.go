package vision

import (
	"github.com/KennethGrace/gracious/modules"
	"github.com/KennethGrace/gracious/modules/util"
)

type Vision struct {
	modules.Module
	feedback util.Feedback
}

func NewVision(vertical int, horizontal int) *Vision {
	return &Vision{}
}

func (v Vision) Begin() {
	//TODO implement me
	panic("implement me")
}
