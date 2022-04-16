package model

// Phenomena are external input to the system. Modules receive Phenomena via the phenomenon registering itself to the
// module. This is done via a call to the given module's implementation of the Module super-class register function.
type Phenomena interface {
	GetQuale() (Quale, error)
}

// Action are external output from the system. Modules send emitted quale to registered Actions.
type Action interface {
	SetQuale(quale Quale) error
}
