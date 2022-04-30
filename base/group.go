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
	newNeurons := make(map[Address]*Neuron)
	// Pass the input pattern through to the output pattern if this group is set for pass-through
	if passThrough {
		firePattern.Composite(main)
	}
	// Test each neuron for firing strength. If no neuron exists to process the signal, make one.
	for featureAddress, feature := range main.Features {
		if neuron, ok := g.neurons[featureAddress]; ok {
			go neuron.Evoke(feature, association, cts, g.LearningControlSignal)
		} else {
			newNeurons[featureAddress] = NewNeuron()
		}
	}
	// Retrieve the firing strength of each neuron and adjust the firing pattern accordingly
	for address, neuron := range g.neurons {
		sum := <-neuron.axon
		firePattern.Features[address] += sum
	}
	// Add the newly needed neurons to the set of neurons in the group
	for newAddress, newNeuron := range newNeurons {
		g.neurons[newAddress] = newNeuron
	}
	firePattern.WinnersTakeAll(0)
	g.firingPattern <- firePattern
}
