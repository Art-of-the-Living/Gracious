package mid

import (
	"github.com/Art-of-the-Living/gracious/base"
)

// A Cluster is a set of Groups such that each Group processes a specific associative DistributedSignal which each
// originates from a different part of the system. When evoked the Cluster aggregates the associational outputs of
// these many associative signals and, after performing a Winner Takes All operation on the aggregate signal, will
// produce a DistributedSignal which, to the system, best represents the composite of associations.
type Cluster struct {
	Id                   string                 //The ID of the system this cluster is a part of.
	PassThrough          bool                   //Determines whether the main signal should be passed through to the output.
	WTA                  bool                   //Determines whether the composite signal should undergo a Winner Takes All operation before leaving the cluster.
	CorrelationThreshold int                    //Determines the necessary threshold for synaptic learning in the cluster.
	groups               map[string]*base.Group //The component groups of this cluster indexed by their linked association.
}

// NewCluster returns a new Cluster.
func NewCluster(id string) *Cluster {
	c := Cluster{Id: id + "-C", groups: make(map[string]*base.Group)}
	return &c
}

// Evoke will test the clusters groups for associational evocations for each associational signal received.
func (c *Cluster) Evoke(main base.DistributedSignal, aSignals ...base.DistributedSignal) base.DistributedSignal {
	var newSig base.DistributedSignal
	// Test each group for firing pattern.
	if c.PassThrough {
		newSig = main
	} else {
		newSig = base.NewDistributedSignal(c.Id)
	}
	// Test the incoming signal for building new neuron groups
	for _, aSig := range aSignals {
		if _, ok := c.groups[aSig.Id]; !ok {
			c.groups[aSig.Id] = base.NewGroup(c.Id + aSig.Id)
		}
	}
	// Slice to hold channels that will be written too
	retrieve := make([]chan base.DistributedSignal, 0, len(aSignals))
	for id, group := range c.groups {
		// For each Group
		for _, associative := range aSignals {
			// For each association
			if associative.Id == id {
				// When the associative ID maps correctly to a Group ID, evoke the neuron group. This keeps us from
				// evoking a group with the wrong association
				go group.Evoke(main, associative, c.CorrelationThreshold)
				retrieve = append(retrieve, group.Pattern)
			}
		}
	}
	// Receive the firing patterns of each group and form a composite signal
	for _, r := range retrieve {
		testPattern := <-r
		newSig.Composite(testPattern)
	}
	// Test for a WTA condition on the composite signal
	if c.WTA {
		newSig.WinnersTakeAll(0)
	}
	return newSig
}
