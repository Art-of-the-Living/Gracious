package base

type Synapse struct {
	// Internal Attributes
	correlationSum int
	weightValue    int
}

func NewSynapse() Synapse {
	syn := Synapse{weightValue: -1, correlationSum: 0}
	return syn
}

//Evoke is invoked at time, T, to train the synapse on the training and feature signals and return the result of
//the associational operation between the feature signal and the synaptic weight. For learning we use an optimized
//correlative Hebbian learning algorithm for training, which prioritizes bit-shifting for multiplications and
//performs +3:-1 incremental steps.
func (syn *Synapse) Evoke(training int, association int, correlation int, learningControl int) int {
	if learningControl > 0 {
		if syn.weightValue != 1 {
			syn.correlationSum += 4 * association * training
			syn.correlationSum -= association
			if syn.correlationSum > correlation {
				syn.weightValue = 1
			}
		}
	}
	return association * syn.weightValue
}
