package model

// Phenomena are external input to the system. Modules receive Phenomena via the phenomenon registering itself to the
// module. This done via a call to the given module's implementation of the Module interfaces register function.
type Phenomena interface {
	GetQuale() (Quale, error)
}
