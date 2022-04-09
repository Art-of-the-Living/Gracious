package base

import (
	"github.com/KennethGrace/gracious/model"
)

// Neuron models the unary behaviour of a single neuron. A neuron is only tangibly useful as a component part of a
// system of Neurons. The goal of each individual neuron is to form an association between synaptic inputs and the
// Neurons "firing" state. The neuron should fire, if and only if, the synaptic inputs
type Neuron struct {
	// Internal Attributes
	synapses []Synapse
}

func NewNeuron(synapseCount int) Neuron {
	synapses := make([]Synapse, synapseCount)
	for i := 0; i < len(synapses); i++ {
		synapses[i] = NewSynapse()
	}
	return Neuron{synapses: synapses}
}

type NeuronProducts struct {
	index int
	value int
}

// GetSumOfWeights returns the amount of synapses which have learnt their weight values. In a bipolar system, this is
// equal to the difference between the number of synapses and the true sum of weights.
func (n *Neuron) GetSumOfWeights() int {
	weightSum := 0
	count := len(n.synapses)
	for i := 0; i < count; i++ {
		weightSum += n.synapses[i].weightValue
	}
	return count + weightSum // Fancy.
}

// Evoke retrieves the strength of the Neuronal signal. The strength is equivalent to the sum of the synaptic
// evocations for a given associative Quale. A synaptic evocation, retrieved with Synapse.Evoke(), triggers synaptic
// learning and returns the product of the bipolar synaptic weight and the feature of the associative Quale.
func (n *Neuron) Evoke(training int, associative model.Quale, correlation int, learningControl int) int {
	sum := 0
	weightSum := 0
	count := len(n.synapses)
	for i := 0; i < count; i++ {
		weightSum += n.synapses[i].weightValue
		feature, _ := associative.GetFeature(i)
		sum += n.synapses[i].Evoke(training, feature, correlation, learningControl)
	}
	return sum
}
