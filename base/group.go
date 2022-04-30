package base

// Group is a set of neurons with a specific associative quale type input and a specific main quale type input and output.
type Group struct {
	// Internal Attributes
	binding               string // The name of the system this neuron group is a part of.
	neurons               map[Address]*Neuron
	LearningControlSignal int
	// Outbound Attributes
	firingPattern chan DistributedSignal
}

func NewGroup(binding string) *Group {
	neurons := make(map[Address]*Neuron)
	ng := Group{neurons: neurons,
		binding:               binding,
		LearningControlSignal: 1,
		firingPattern:         make(chan DistributedSignal),
	}
	return &ng
}

func (g *Group) Evoke(main DistributedSignal, association DistributedSignal, passThrough bool, cts int) {
	firePattern := NewDistributedSignal(g.binding + ":group")
	// Pass the input pattern through to the output pattern if this group is set for pass-through
	if passThrough {
		firePattern.Composite(main)
	}
	// Test the incoming signal for building new neurons
	for addr := range main.Features {
		if _, ok := g.neurons[addr]; !ok {
			g.neurons[addr] = NewNeuron()
		}
	}
	// Test each neuron for firing strength.
	for addr, neuron := range g.neurons {
		go neuron.Evoke(main.Features[addr], association, cts, g.LearningControlSignal)
	}
	// Retrieve the firing strength of each neuron and adjust the firing pattern accordingly
	for address, neuron := range g.neurons {
		sum := <-neuron.axon
		firePattern.Features[address] += sum
	}
	firePattern.WinnersTakeAll(0)
	g.firingPattern <- firePattern
}
