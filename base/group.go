package base

import (
	"github.com/KennethGrace/gracious/model"
)

// Group is a set of neurons with a specific associative quale type input and a specific main quale type input and output.
type Group struct {
	// Internal Attributes
	binding string // The name of the system this neuron group is a part of.
	neurons map[model.Address]Neuron
	// Inbound Attributes
	CorrelationThresholdSignal int
	LearningControlSignal      int
	Association                model.Quale
	PassThrough                bool
	// Outbound Attributes
	Match   int
	Novelty int
}

func NewGroup(binding string) *Group {
	neurons := make(map[model.Address]Neuron)
	ng := Group{neurons: neurons, binding: binding, LearningControlSignal: 1}
	return &ng
}

func (g *Group) SetAssociation(a model.Quale) {
	g.Association = a
}

// Evoke updates the Neuron Group for the moment of time, T.
func (g *Group) Evoke(main model.Quale) model.Quale {
	newQuale := model.NewQuale()
	if g.PassThrough {
		newQuale = main
	}
	for featureAddress, feature := range main.GetFeatures() {
		if neuron, ok := g.neurons[featureAddress]; ok {
			sum := neuron.Evoke(feature, g.Association, g.CorrelationThresholdSignal, g.LearningControlSignal)
			newQuale.AdjustFeature(featureAddress, sum)
		} else {
			g.neurons[featureAddress] = NewNeuron()
		}
	}
	newQuale.WinnersTakeAll(0)
	return newQuale
}
