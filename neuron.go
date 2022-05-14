package gracious

import "sync"

// Neuron models the unary behaviour of a single neuron. A neuron is only tangibly useful as a component part of a
// system of Neurons. The goal of each individual neuron is to form an association between synaptic inputs and the
// Neurons "firing" state. The neuron should fire, if and only if, the synaptic inputs
type Neuron struct {
	// Internal Attributes
	synapses        map[Address]*Synapse
	axon            int
	match           bool
	novelty         bool
	learningEnabled bool
	wg              sync.WaitGroup
}

func NewNeuron() *Neuron {
	synapses := make(map[Address]*Synapse)
	n := Neuron{synapses: synapses, learningEnabled: true}
	return &n
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
func (n *Neuron) Evoke(training int, associative QualitativeSignal, correlation int) {
	sum := 0
	// Test the neuron synaptic associative evocations, if there is not a synapse present to handle the association
	// feature then a new synapse will be made.
	for featureAddress, feature := range associative.Features {
		if syn, ok := n.synapses[featureAddress]; ok {
			value := syn.Evoke(feature)
			sum += value
		} else {
			n.synapses[featureAddress] = NewSynapse()
		}
	}
	// Training should occur on the condition of a novelty state being produced by
	// the current system and only when learning has been enabled
	if n.learningEnabled {
		for featureAddress, feature := range associative.Features {
			if syn, ok := n.synapses[featureAddress]; ok {
				if (sum <= 0) && (training != 0) {
					n.novelty = true
					syn.Train(training, feature, correlation)
				}
			}
		}
	}
	// In the case that both signals are the same polarity, match is true.
	// In the case that both signals are of different polarity, match is false.
	n.match = ((sum > 0) && (training > 0)) || ((sum <= 0) && (training <= 0))
	n.axon = sum
}

// AsyncEvoke will Evoke this Neuron as a member of a WaitGroup
func (n *Neuron) AsyncEvoke(training int, associative QualitativeSignal, correlation int, wg *sync.WaitGroup) {
	defer wg.Done()
	n.Evoke(training, associative, correlation)
}
