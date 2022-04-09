package base

// Neuron models the unary behaviour of a single neuron. A neuron may be of two types, Basic and Broadcast.
type Neuron struct {
	Synapses                   []*Synapse
	MainSignal                 *int
	CorrelationThresholdSignal *int
	LearningControlSignal      *int
	SumOfWeights               int // A value indicating the number of weights over the weight minimum.
	SumOfSynapticFirings       int
	index                      int // A value denoting the neurons position within a NeuronGroup
}

func NewNeuron(index int, synapseCount int) *Neuron {
	synapses := make([]*Synapse, synapseCount)
	for i := 0; i < len(synapses); i++ {
		synapses[i] = NewSynapse()
	}
	n := Neuron{Synapses: synapses, index: index}
	return &n
}

type NeuronProducts struct {
	index int
	value int
}

// Evoke retrieves the strength of the Neuron signal strength, but only after computing a new strength value
// based on the current state of the Neuron's component synapses. If you want to retrieve the Neuron's signal strength
// without evocation of the synapses or testing for learning use "Evoke" directly.
func (n *Neuron) Evoke() {
	sum := 0
	weightSum := 0
	count := len(n.Synapses)
	for i := 0; i < count; i++ {
		weightSum += n.Synapses[i].weightValue
		sum += n.Synapses[i].Evoke(*n.MainSignal, *n.CorrelationThresholdSignal)
	}
	//println("Neuron", &n, "Evoked: ", sum)
	n.SumOfWeights = count + weightSum
	n.SumOfSynapticFirings = sum
}
