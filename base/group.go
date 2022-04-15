package base

import "github.com/KennethGrace/gracious/model"

// Group is a set of neurons with a specific associative quale type input and a specific main quale type input and output.
type Group struct {
	// Internal Attributes
	binding string // The name of the system this neuron group is a part of.
	neurons map[model.Address]Neuron
	// Inbound Attributes
	CorrelationThresholdSignal int
	LearningControlSignal      int
	// Outbound Attributes
	Match   int
	Novelty int
}

func NewGroup(binding string) *Group {
	neurons := make(map[model.Address]Neuron)
	ng := Group{neurons: neurons, binding: binding}
	return &ng
}

// Evoke updates the Neuron Group for the moment of time, T.
func (ng *Group) Evoke(main model.Quale, association model.Quale) model.Quale {
	sigMax := 0
	newQuale := model.NewQuale()
	for featureAddress, feature := range main.GetFeatures() {
		if neuron, ok := ng.neurons[featureAddress]; ok {
			sum := neuron.Evoke(feature, association, ng.CorrelationThresholdSignal, ng.LearningControlSignal)
			if sum > sigMax {
				sigMax = sum
			}
			_ = newQuale.SetFeature(featureAddress, sum)
		} else {
			ng.neurons[featureAddress] = NewNeuron()
		}
	}
	for address, value := range newQuale.GetFeatures() {
		if value >= sigMax {
			_ = newQuale.SetFeature(address, 1)
		} else {
			_ = newQuale.SetFeature(address, 0)
		}
	}
	return newQuale
}
