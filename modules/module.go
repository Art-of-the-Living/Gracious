package modules

import "github.com/KennethGrace/gracious/model"

// A Module is a qualar processing system which produces qualia based on associative evocations from other qualia. In
// some cases a module is connected externally via phenomena. Phenomena produce phenomenal qualia in the system.
type Module struct {
	dispatch  *model.Dispatch
	Phenomena model.Phenomena
}

func (m *Module) Begin() {
}

func (m *Module) SetDispatch(dispatch *model.Dispatch) {
	m.dispatch = dispatch
}

func (m *Module) Publish(q model.Quale) {
	m.dispatch.Distribute(q)
}
