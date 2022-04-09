package base

type Synapse struct {
	Value                 int
	AssociativeSignal     *int
	LearningControlSignal *int
	correlationSum        int
	weightValue           int
}

func NewSynapse() *Synapse {
	syn := Synapse{weightValue: -1, correlationSum: 0}
	return &syn
}

// Evoke is invoked at time, T, to train the synapse on the active signal values and return the result of
// the associational operation between the resulting signal values according to the synaptic weight. Here we use
// a bipolar correlative Hebbian learning algorithm.
func (syn *Synapse) Evoke(sValue int, correlationThreshold int) int {
	if *syn.LearningControlSignal > 0 {
		if syn.weightValue != 1 {
			a := *syn.AssociativeSignal
			s := sValue
			syn.correlationSum += 4 * a * s
			syn.correlationSum -= a
			if syn.correlationSum > correlationThreshold {
				syn.weightValue = 1
			}
		}
	}
	syn.Value = syn.weightValue * *syn.AssociativeSignal
	//fmt.Println("Synapse", &syn, "Evoked:", syn.Value, "via", syn.weightValue, "x", *syn.AssociativeSignal)
	return syn.Value
}
