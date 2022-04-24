package base

// Cluster is a set of Groups such that each Group processes a specific associative quale type, but all Groups receive
// the same main quale type and quale. When evoked a cluster returns only the strongest quale. A cluster ensures that
// quale are evoked associatively only by the proper quale type for the association.
type Cluster struct {
	binding      string            //The name of the system this cluster is a part of.
	groups       map[string]*Group //The component groups of this cluster indexed by their linked association.
	Destinations map[string]chan Quale
	Main         chan Quale
	Active       bool
}

func NewCluster(binding string) *Cluster {
	c := Cluster{binding: binding, groups: make(map[string]*Group)}
	c.Destinations = make(map[string]chan Quale)
	return &c
}

func (c *Cluster) AddDestination(binding string, channel chan Quale) {
	c.Destinations[binding] = channel
}

func (c *Cluster) AddNewGroup(binding string) *Group {
	ng := NewGroup(c.binding)
	c.groups[binding] = ng
	return ng
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

func (c *Cluster) Evoke() {
	strongestQuale := NewQuale()
	main := <-c.Main
	out := make(map[string]chan Quale, len(c.groups))
	for binding, group := range c.groups {
		out[binding] = make(chan Quale)
		go group.Evoke(main, out[binding])
	}
	for binding, _ := range c.groups {
		out := <-out[binding]
		if out.Strength() > strongestQuale.Strength() {
			strongestQuale = out
		}
	}
	for _, destination := range c.Destinations {
		select {
		case destination <- strongestQuale:
		}
	}
}
