package components

import (
	"github.com/Art-of-the-Living/gracious/base"
	"github.com/Art-of-the-Living/gracious/mid"
)

// A Component is a simple structural unit in an architecture. A component performs a specific type of role in an
// architecture, such as comparing and compositing associations, enhancing feedback, boosting and decaying
// signal values, distributing signals across the network, controlling motor systems, and processing sensory
// information. Component types are named by the produced distributed signals implied meaning.
type Component interface {
	SetDistributedSignals(signal base.DistributedSignal, net *mid.SubNet) // Sets the Signal and Network for this component.
	Evoke() base.DistributedSignal                                        // Evoking a component returns a DistributedSignal.
	GetName() string                                                      // Every component should have a name. It should be as unique as possible.
}
