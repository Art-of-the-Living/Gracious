package base

// Group is a set of neurons with a specific associative quale type input and a specific main quale type input and output.
type Group struct {
	// Internal Attributes
	binding string // The name of the system this neuron group is a part of.
	neurons map[Address]*Neuron
	// Inbound Attributes
	CorrelationThresholdSignal int
	LearningControlSignal      int
	PassThrough                bool
	// Outbound Attributes
	firingPattern chan DistributedSignal
	Match         int
	Novelty       int
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

func (g *Group) Evoke(main DistributedSignal, association DistributedSignal) {
	firePattern := NewQuale()
	newNeurons := make(map[Address]*Neuron)
	// Pass the input pattern through to the output pattern if this group is set for pass-through
	if g.PassThrough {
		firePattern = main
	}
	// Test each neuron for firing strength. If no neuron exists to process the signal, make one.
	for featureAddress, feature := range main.GetFeatures() {
		if neuron, ok := g.neurons[featureAddress]; ok {
			go neuron.Evoke(feature, association, g.CorrelationThresholdSignal, g.LearningControlSignal)
		} else {
			newNeurons[featureAddress] = NewNeuron()
		}
	}
	// Retrieve the firing strength of each neuron and adjust the firing pattern accordingly
	for address, neuron := range g.neurons {
		sum := <-neuron.axon
		firePattern.AdjustFeature(address, sum)
	}
	// Add the newly needed neurons to the set of neurons in the group
	for newAddress, newNeuron := range newNeurons {
		g.neurons[newAddress] = newNeuron
	}
	firePattern.WinnersTakeAll(0)
	g.firingPattern <- firePattern
}
