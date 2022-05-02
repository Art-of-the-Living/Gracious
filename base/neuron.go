package base

// Neuron models the unary behaviour of a single neuron. A neuron is only tangibly useful as a component part of a
// system of Neurons. The goal of each individual neuron is to form an association between synaptic inputs and the
// Neurons "firing" state. The neuron should fire, if and only if, the synaptic inputs
type Neuron struct {
	// Internal Attributes
	synapses map[Address]*Synapse
	axon     chan int
}

func NewNeuron() *Neuron {
	synapses := make(map[Address]*Synapse)
	return &Neuron{synapses: synapses, axon: make(chan int)}
}

// GetSumOfWeights returns the amount of synapses which have learnt their weight values. In a bipolar system, this is
// equal to the difference between the number of synapses and the true sum of weights.
func (n *Neuron) GetSumOfWeights() int {
	weightSum := 0
	count := len(n.synapses)
	for _, syn := range n.synapses {
		weightSum += syn.weightValue
	}
	return count + weightSum // Fancy.
}

// Evoke tests the neuron for firing and writes the fired value to the 'axon'
// channel. If the firing state does not evoke in the presence of the training
// signal, the synaptic association trains itself.
func (n *Neuron) Evoke(training int, associative DistributedSignal, correlation int, learningControl int) {
	sum := 0
	// Test the neuron synaptic associative evocations
	for featureAddress, feature := range associative.Features {
		if syn, ok := n.synapses[featureAddress]; ok {
			value := syn.Evoke(feature)
			sum += value
		} else {
			n.synapses[featureAddress] = NewSynapse()
		}
	}
	// Training should occur on the condition of a novelty state being produced by
	// the current system
	for featureAddress, feature := range associative.Features {
		if syn, ok := n.synapses[featureAddress]; ok {
			if (sum < 0) && (training != 0) {
				syn.Train(training, feature, correlation)
			}
		}
	}
	n.axon <- sum
}
