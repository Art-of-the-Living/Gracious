package learners

// The Synapse performs the crucial job of connecting associations to neuron groups. Each synapse has a certain weight
// which is either negative 1 or positive one. Once the weight has been set to one via +3:-1 Hebbian learning, the weight
// must be reset by a raise in the correlation sum threshold value.
type Synapse struct {
	// Internal Attributes
	correlationSum int
	weightValue    int
}

// NewSynapse initializes a new Synapse with a -1 weight value and a 0 correlation sum. A pointer to the Synapse is
// returned.
func NewSynapse() *Synapse {
	syn := &Synapse{weightValue: -1, correlationSum: 0}
	return syn
}

// Evoke is invoked at time, T, to return the result of the associational
// operation between the feature signal and the synaptic weight.
func (syn *Synapse) Evoke(association int) int {
	return association * syn.weightValue
}

// Train trains the synapse for later evocation. For learning we use an optimized
// correlative Hebbian learning algorithm for training, which prioritizes
// bit-shifting for multiplications and performs +3:-1 incremental steps.
func (syn *Synapse) Train(training int, association int, correlation int) {
	if syn.correlationSum > correlation {
		syn.weightValue = 1
	} else {
		syn.weightValue = -1
		syn.correlationSum += 4 * association * training
		syn.correlationSum -= association
	}
}
