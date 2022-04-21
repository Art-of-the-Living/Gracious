package modules

import (
	"github.com/KennethGrace/gracious/base"
	"github.com/KennethGrace/gracious/model"
)

// A Module is a qualar processing system which produces qualia based on associative evocations from other qualia. In
// some cases a module is connected externally via phenomena. Phenomena produce phenomenal qualia in the system.
type Module struct {
	Dispatch  *base.Dispatch
	Phenomena model.Phenomena
	Action    model.Action
}

func (m *Module) Begin(time int) {
}

func (m *Module) SetDispatch(dispatch *base.Dispatch) {
	m.Dispatch = dispatch
}

func (m *Module) GetDispatch() *base.Dispatch {
	return m.Dispatch
}

func (m *Module) Publish(q model.Quale) {
	m.Dispatch.Distribute(q)
}
