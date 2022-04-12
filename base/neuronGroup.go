package base

import "github.com/KennethGrace/gracious/model"

type NeuronGroup struct {
	// Internal Attributes
	neurons map[model.Address]Neuron
	// Inbound Attributes
	main                       *model.Quale //At initialization a new Quale is made.
	association                *model.Quale //At initialization a new Quale is made.
	CorrelationThresholdSignal *int
	LearningControlSignal      *int
	// Outbound Attributes
	Match   int
	Novelty int
	MainOut model.Quale
}

func NewNeuronGroup() *NeuronGroup {
	neurons := make(map[model.Address]Neuron)
	main := model.NewQuale()
	associative := model.NewQuale()
	correlationThresholdSignal := new(int)
	learningControlSignal := new(int)
	ng := NeuronGroup{neurons: neurons}
	ng.main = &main
	ng.association = &associative
	ng.CorrelationThresholdSignal = correlationThresholdSignal
	ng.LearningControlSignal = learningControlSignal
	ng.MainOut = model.NewQuale()
	return &ng
}

func (ng *NeuronGroup) SetMain(q *model.Quale) {
	ng.main = q
}

func (ng *NeuronGroup) SetAssociation(q *model.Quale) {
	ng.association = q
}

// Evoke updates the Neuron Group for the moment of time, T.
func (ng *NeuronGroup) Evoke() {
	sigMax := 0
	rawQuale := model.NewQuale()
	for featureAddress, feature := range ng.main.GetFeatures() {
		if neuron, ok := ng.neurons[featureAddress]; ok {
			sum := neuron.Evoke(feature, *ng.association, *ng.CorrelationThresholdSignal, *ng.LearningControlSignal)
			if sum > sigMax {
				sigMax = sum
			}
			_ = rawQuale.SetFeature(featureAddress, sum)
		} else {
			ng.neurons[featureAddress] = NewNeuron()
		}
	}
	ng.MainOut.Zero()
	for address, value := range rawQuale.GetFeatures() {
		if value >= sigMax {
			_ = ng.MainOut.SetFeature(address, 1)
		} else {
			_ = ng.MainOut.SetFeature(address, 0)
		}
	}
}
