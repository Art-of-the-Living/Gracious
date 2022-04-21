package base

import "github.com/KennethGrace/gracious/model"

// AutoAssociativeMemory (AAM) is used to store Quale in memory temporally. An AAM is defined by looping the main
// model.Quale back into the associative model.Quale. Therefore, the reappearance of the whole or a sufficient part of
// the main signal can reproduce the whole main signal.
type AutoAssociativeMemory struct {
	Group
}

func NewAutoAssociativeMemory(binding string) *AutoAssociativeMemory {
	return &AutoAssociativeMemory{Group: *NewGroup(binding)}
}

// Evoke updates the Neuron Group for the moment of time, T.
func (aam *AutoAssociativeMemory) Evoke(main model.Quale) model.Quale {
	sigMax := 0
	newQuale := model.NewQuale()
	if aam.PassThrough {
		newQuale = main
	}
	for featureAddress, feature := range main.GetFeatures() {
		if neuron, ok := aam.neurons[featureAddress]; ok {
			sum := neuron.Evoke(feature, main, aam.CorrelationThresholdSignal, aam.LearningControlSignal)
			if sum > sigMax {
				sigMax = sum
			}
			newQuale.SetFeature(featureAddress, sum)
		} else {
			aam.neurons[featureAddress] = NewNeuron()
		}
	}
	for address, value := range newQuale.GetFeatures() {
		if value < sigMax {
			newQuale.SetFeature(address, 0)
		}
	}
	return newQuale
}
