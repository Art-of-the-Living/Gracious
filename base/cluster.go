package base

// Cluster is a set of Groups such that each Group processes a specific associative quale type, but all Groups receive
// the same main quale type and quale. When evoked a cluster returns only the strongest quale. A cluster ensures that
// quale are evoked associatively only by the proper quale type for the association.
type Cluster struct {
	binding      string            //The name of the system this cluster is a part of.
	groups       map[string]*Group //The component groups of this cluster indexed by their linked association.
	Destinations map[string]chan DistributedSignal
	Main         chan DistributedSignal
	Active       bool
}

func NewCluster(binding string) *Cluster {
	c := Cluster{binding: binding, groups: make(map[string]*Group)}
	c.Destinations = make(map[string]chan DistributedSignal)
	return &c
}

func (c *Cluster) SetPassThrough(pass bool) {
	for _, group := range c.groups {
		group.PassThrough = pass
	}
}

func (c *Cluster) SetCorrelationThreshold(signal int) {
	for _, group := range c.groups {
		group.CorrelationThresholdSignal = signal
	}
}

// Evoke will test the clusters groups for possible signals.
func (c *Cluster) Evoke(main DistributedSignal, associates map[string]DistributedSignal) DistributedSignal {
	strongestSignal := NewDistributedSignal()
	newGroups := make(map[string]*Group)
	// Test each group for firing patterns. If no group exists to handle the associated signal, create one.
	for binding, signal := range associates {
		if group, ok := c.groups[binding]; ok {
			go group.Evoke(main, signal)
		} else {
			newGroups[binding] = NewGroup(binding)
		}
	}
	// Receive the firing patterns of each group and
	for _, group := range c.groups {
		testPattern := <-group.firingPattern
		if testPattern.Strength() > strongestSignal.Strength() {
			strongestSignal = testPattern
		}
	}
	// Add any newly created groups to the cluster groups
	for newBinding, newGroup := range newGroups {
		c.groups[newBinding] = newGroup
	}
	return strongestSignal
}
