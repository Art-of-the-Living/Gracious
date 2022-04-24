package base

// Group is a set of neurons with a specific associative quale type input and a specific main quale type input and output.
type Group struct {
	// Internal Attributes
	binding string // The name of the system this neuron group is a part of.
	neurons map[Address]Neuron
	// Inbound Attributes
	CorrelationThresholdSignal int
	LearningControlSignal      int
	Association                chan Quale
	PassThrough                bool
	// Outbound Attributes
	Match   int
	Novelty int
}

func NewGroup(binding string) *Group {
	neurons := make(map[Address]Neuron)
	ng := Group{neurons: neurons, binding: binding, LearningControlSignal: 1}
	return &ng
}

func (g *Group) Evoke(main Quale, out chan Quale) {
	association := NewQuale()
	select {
	case association = <-g.Association:
	}
	newQuale := NewQuale()
	if g.PassThrough {
		newQuale = main
	}
	for featureAddress, feature := range main.GetFeatures() {
		if neuron, ok := g.neurons[featureAddress]; ok {
			sum := neuron.Evoke(feature, association, g.CorrelationThresholdSignal, g.LearningControlSignal)
			newQuale.AdjustFeature(featureAddress, sum)
		} else {
			g.neurons[featureAddress] = NewNeuron()
		}
	}
	newQuale.WinnersTakeAll(0)
	out <- newQuale
}
