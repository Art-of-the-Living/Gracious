package base

type NeuronGroup struct {
	Neurons                    []*Neuron
	MainSignalLines            []*int
	AssociativeSignalLines     []*int
	CorrelationThresholdSignal *int
	LearningControlSignal      *int
	Match                      int
	Novelty                    int
	OutSignal                  []int
}

func NewNeuronGroup(neuronCount int, synapseCount int) *NeuronGroup {
	neurons := make([]*Neuron, neuronCount)
	mainSignalLines := make([]*int, neuronCount)
	associativeSignalLines := make([]*int, synapseCount)
	correlationThresholdSignal := new(int)
	learningControlSignal := new(int)
	for j := 0; j < len(associativeSignalLines); j++ {
		associativeSignalLines[j] = new(int)
	}
	for i := 0; i < len(neurons); i++ {
		neurons[i] = NewNeuron(i, synapseCount)
		neurons[i].MainSignal = new(int)
		mainSignalLines[i] = neurons[i].MainSignal
		neurons[i].LearningControlSignal = learningControlSignal
		neurons[i].CorrelationThresholdSignal = correlationThresholdSignal
		for j := 0; j < len(neurons[i].Synapses); j++ {
			neurons[i].Synapses[j].AssociativeSignal = associativeSignalLines[j]
			neurons[i].Synapses[j].LearningControlSignal = learningControlSignal
		}
	}
	ng := NeuronGroup{Neurons: neurons}
	ng.MainSignalLines = mainSignalLines
	ng.AssociativeSignalLines = associativeSignalLines
	ng.CorrelationThresholdSignal = correlationThresholdSignal
	ng.LearningControlSignal = learningControlSignal
	ng.OutSignal = make([]int, neuronCount)
	return &ng
}

func (ng *NeuronGroup) SetNewMainSignalLines(signals []*int) {
	if len(ng.Neurons) != len(signals) {
		panic("Length of new signals not equal to neuron count")
	}
	for i := 0; i < len(ng.Neurons); i++ {
		ng.Neurons[i].MainSignal = signals[i]
	}
	ng.MainSignalLines = signals
}

func (ng *NeuronGroup) SetNewAssociativeSignalLines(signals []*int) {
	for i := 0; i < len(ng.Neurons); i++ {
		if len(ng.Neurons[i].Synapses) != len(signals) {
			panic("Length of new signals not equal to synapse count")
		}
		for j := 0; i < len(ng.Neurons[i].Synapses); j++ {
			ng.Neurons[i].Synapses[j].AssociativeSignal = signals[j]
		}
	}
	ng.AssociativeSignalLines = signals
}

// Evoke updates the Neuron Group for the moment of time, T.
func (ng *NeuronGroup) Evoke() {
	sigMax := 0
	sumLines := make([]int, len(ng.Neurons))
	for i := 0; i < len(ng.Neurons); i++ {
		ng.Neurons[i].Evoke()
		if ng.Neurons[i].SumOfSynapticFirings > sigMax {
			sigMax = ng.Neurons[i].SumOfSynapticFirings
		}
		sumLines[i] = ng.Neurons[i].SumOfSynapticFirings
	}
	for i := 0; i < len(sumLines); i++ {
		if sumLines[i] < sigMax {
			ng.OutSignal[i] = 0
		} else {
			ng.OutSignal[i] = 1
		}
	}
}
