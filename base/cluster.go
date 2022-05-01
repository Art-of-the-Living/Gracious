package base

// Cluster is a set of Groups such that each Group processes a specific associative DistributedSignal which each
// originates from a different part of the system. When evoked the Cluster aggregates the associational outputs of
// these many associative signals and, after performing a Winner Takes All operation on the aggregate signal, will
// produce a DistributedSignal which, to the system, best represents the composite of associations.
type Cluster struct {
	binding              string            //The name of the system this cluster is a part of.
	groups               map[string]*Group //The component groups of this cluster indexed by their linked association.
	PassThrough          bool
	CorrelationThreshold int
}

// NewCluster returns a new Cluster
func NewCluster(binding string) *Cluster {
	c := Cluster{binding: binding, groups: make(map[string]*Group)}
	return &c
}

// Evoke will test the clusters groups for possible signals.
func (c *Cluster) Evoke(main DistributedSignal, associates map[string]DistributedSignal) DistributedSignal {
	newDistributedSignal := NewDistributedSignal(c.binding + ":evocation")
	// Test the incoming signal for building new neuron groups
	for binding := range associates {
		if _, ok := c.groups[binding]; !ok {
			c.groups[binding] = NewGroup(binding)
		}
	}
	// Test each group for firing pattern.
	for binding, group := range c.groups {
		go group.Evoke(main, associates[binding], c.PassThrough, c.CorrelationThreshold)
	}
	// Receive the firing patterns of each group and
	for _, group := range c.groups {
		testPattern := <-group.firingPattern
		newDistributedSignal.Composite(testPattern)
	}
	return newDistributedSignal
}
