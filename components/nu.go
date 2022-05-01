package components

import "github.com/Art-of-the-Living/gracious/base"

// Nu (𝛎) is phenomenal experience as direct sensory observations. A Nu component processes incoming raw-data
// into distributed signals that the system can only then begin to work with. Unlike other components a Nu component
// doesn't have any DistributedSignal inputs, instead Nu produces a DistributedSignal based on the output of an
// externally implemented function.
type Nu struct {
	name      string
	processor func() base.DistributedSignal
}

// NewNu creates a new Nu, haha
func NewNu(name string) *Nu {
	return &Nu{name: "Nu:" + name}
}

// GetName returns the name of this Nu
func (n *Nu) GetName() string {
	return n.name
}

// SetProcessor sets the function that should be run when this Nu component is evoked.
func (n *Nu) SetProcessor(processor func() base.DistributedSignal) {
	n.processor = processor
}

// Evoke returns the result of the specified processor function. If no processor is defined, then an empty signal will
// be returned.
func (n *Nu) Evoke() base.DistributedSignal {
	if n.processor != nil {
		return n.processor()
	} else {
		return base.NewDistributedSignal(n.name + ":voidPhenomena")
	}
}
