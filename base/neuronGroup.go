package base

import "github.com/KennethGrace/gracious/model"

type NeuronGroup struct {
	// Internal Attributes
	neurons []Neuron
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

func NewNeuronGroup(neuronCount int, synapseCount int) *NeuronGroup {
	neurons := make([]Neuron, neuronCount)
	main := model.NewQuale(neuronCount)
	associative := model.NewQuale(synapseCount)
	correlationThresholdSignal := new(int)
	learningControlSignal := new(int)
	for i := 0; i < len(neurons); i++ {
		neurons[i] = NewNeuron(synapseCount)
		for j := 0; j < len(neurons[i].synapses); j++ {
		}
	}
	ng := NeuronGroup{neurons: neurons}
	ng.main = &main
	ng.association = &associative
	ng.CorrelationThresholdSignal = correlationThresholdSignal
	ng.LearningControlSignal = learningControlSignal
	ng.MainOut = model.NewQuale(neuronCount)
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
	sumLines := make([]int, len(ng.neurons))
	for i := 0; i < len(ng.neurons); i++ {
		training, _ := ng.main.GetFeature(i)
		sum := ng.neurons[i].Evoke(training, *ng.association, *ng.CorrelationThresholdSignal, *ng.LearningControlSignal)
		if sum > sigMax {
			sigMax = sum
		}
		sumLines[i] = sum
	}
	for i := 0; i < len(sumLines); i++ {
		if sumLines[i] < sigMax {
			_ = ng.MainOut.SetFeature(i, 0)
		} else {
			_ = ng.MainOut.SetFeature(i, 1)
		}
	}
}
