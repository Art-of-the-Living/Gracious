package base

// Group is a set of neurons with a specific associative DistributedSignal type input and a specific main
// DistributedSignal type input and output.
type Group struct {
	Id string // The name of the system this neuron group is a part of.
	// Internal Attributes
	neurons map[Address]*Neuron
	// Outbound Attributes
	pattern chan DistributedSignal
}

func NewGroup(id string) *Group {
	neurons := make(map[Address]*Neuron)
	ng := Group{neurons: neurons,
		Id:      id,
		pattern: make(chan DistributedSignal),
	}
	return &ng
}

func (g *Group) GetFiringPattern() DistributedSignal {
	return <-g.pattern
}

func (g *Group) Evoke(main DistributedSignal, association DistributedSignal, passThrough bool, cts int) {
	firePattern := NewDistributedSignal(g.Id + ":group")
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
		go neuron.Evoke(main.Features[addr], association, cts)
	}
	// Retrieve the firing strength of each neuron and adjust the firing pattern accordingly
	for address, neuron := range g.neurons {
		sum := <-neuron.axon
		firePattern.Features[address] += sum
	}
	firePattern.WinnersTakeAll(0)
	g.pattern <- firePattern
}
