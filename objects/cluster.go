package objects

import (
	"github.com/Art-of-the-Living/gracious/base"
)

// A Cluster is a set of Groups such that each Group processes a specific associative DistributedSignal which each
// originates from a different part of the system. When evoked the Cluster aggregates the associational outputs of
// these many associative signals and, after performing a Winner Takes All operation on the aggregate signal, will
// produce a DistributedSignal which, to the system, best represents the composite of associations.
//
// When a cluster is called to process a slice of main signals then the SAME Group instances will process each
// DistributedSignal serially. The product will be a slice of the same length containing the normal product
// of each associational evocation. Be advised that learning will occur on each DistributedSignal. For parallel
// slice with unique learning for each temporal signal, use a SuperCluster.
type Cluster struct {
	Id                   string                 //The name of the system this cluster is a part of.
	PassThrough          bool                   //Determines whether the main signal should be passed through to the output.
	WTA                  bool                   //Determines whether the composite signal should undergo a Winner Takes All operation before leaving the cluster.
	CorrelationThreshold int                    //Determines the necessary threshold for synaptic learning in the cluster.
	groups               map[string]*base.Group //The component groups of this cluster indexed by their linked association.
}

// NewCluster returns a new Cluster.
func NewCluster(id string) *Cluster {
	c := Cluster{Id: id, groups: make(map[string]*base.Group)}
	return &c
}

// Evoke will test the clusters groups for associational evocations for each main signal received.
func (c *Cluster) Evoke(aSignals []base.DistributedSignal, mSignals ...base.DistributedSignal) []base.DistributedSignal {
	var newDistributedSignals []base.DistributedSignal
	// Test for PassThrough or create a fresh slice of DistributedSignals.
	if c.PassThrough {
		newDistributedSignals = mSignals
	} else {
		newDistributedSignals = make([]base.DistributedSignal, len(mSignals))
		for mi, mSig := range mSignals {
			newDistributedSignals[mi] = base.NewDistributedSignal(mSig.Id)
		}
	}
	// Test the incoming signal for building new neuron groups
	for _, aSig := range aSignals {
		if _, ok := c.groups[aSig.Id]; !ok {
			c.groups[aSig.Id] = base.NewGroup(aSig.Id)
		}
	}
	// Test each group for firing pattern.
	for mi, main := range mSignals { // For each Main
		retrieve := make([]func() base.DistributedSignal, len(aSignals))
		for id, group := range c.groups { // For each Group
			for ai, aSig := range aSignals { // For each A Signal
				if aSig.Id == id { // When the aSig ID maps correctly to a Group ID
					go group.Evoke(main, aSig, c.PassThrough, c.CorrelationThreshold)
					retrieve[ai] = group.GetFiringPattern
				}
			}
		}
		// Receive the firing patterns of each group and form a composite signal
		for _, r := range retrieve {
			testPattern := r()
			newDistributedSignals[mi].Composite(testPattern)
		}
	}
	return newDistributedSignals
}
