package base

import "github.com/KennethGrace/gracious/model"

// AutoAssociativeMemory (AAM) is used to store Quale in memory temporally. An AAM is defined by looping the main
// model.Quale back into the associative model.Quale. Therefore, a change to either the Association Quale or the
// Main Quale will result in a change to both.
type AutoAssociativeMemory struct {
	NeuronGroup
}

func NewAutoAssociativeMemory(neuronCount int, synapseCount int) *AutoAssociativeMemory {
	aam := AutoAssociativeMemory{}
	aam.NeuronGroup = *NewNeuronGroup(neuronCount, synapseCount)
	aam.association = aam.main
	return &aam
}

func (aam *AutoAssociativeMemory) SetMain(q *model.Quale) {
	aam.association = q
	aam.main = q
}

func (aam *AutoAssociativeMemory) SetAssociation(q *model.Quale) {
	aam.association = q
	aam.main = q
}
