package base

import "github.com/KennethGrace/gracious/model"

// Cluster is a set of Groups such that each Group processes a specific associative quale type, but all Groups receive
// the same main quale type and quale. When evoked a cluster returns only the strongest quale. A cluster ensures that
// quale are evoked associatively only by the proper quale type for the association.
type Cluster struct {
	binding string            //The name of the system this cluster is a part of.
	groups  map[string]*Group //The component groups of this cluster indexed by their linked association.
}

func NewCluster(binding string) *Cluster {
	c := Cluster{binding: binding, groups: make(map[string]*Group)}
	return &c
}

func (c *Cluster) AddNewGroup(binding string) *Group {
	ng := NewGroup(c.binding)
	c.groups[binding] = ng
	return ng
}

func (c *Cluster) AddNewGroups(bindings []string) []*Group {
	groups := make([]*Group, len(bindings))
	for i, binding := range bindings {
		ng := NewGroup(c.binding)
		c.groups[binding] = ng
		groups[i] = ng
	}
	return groups
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

func (c *Cluster) Evoke(main model.Quale) model.Quale {
	strongestQuale := model.NewQuale()
	for _, group := range c.groups {
		q := group.Evoke(main)
		if q.Strength() > strongestQuale.Strength() {
			strongestQuale = q
		}
	}
	return strongestQuale
}
