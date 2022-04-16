package base

import "github.com/KennethGrace/gracious/model"

// Cluster is a set of Groups such that each Group processes a specific associative quale type, but all Groups receive
// the same main quale type and quale. When evoked a cluster returns only the strongest quale. A cluster ensures that
// quale are evoked associatively only by the proper quale type for the association.
type Cluster struct {
	binding string            //The name of the system this cluster is a part of.
	groups  map[string]*Group //The component groups of this cluster.
}

func NewCluster(binding string) *Cluster {
	c := Cluster{binding: binding, groups: make(map[string]*Group)}
	return &c
}

func (c *Cluster) AddNewGroup(binding string) {
	ng := NewGroup(c.binding)
	ng.LearningControlSignal = 0
	ng.CorrelationThresholdSignal = 0
	c.groups[binding] = ng
}

func (c *Cluster) AddNewGroups(bindings []string) {
	for _, binding := range bindings {
		ng := NewGroup(c.binding)
		ng.LearningControlSignal = 0
		ng.CorrelationThresholdSignal = 0
		c.groups[binding] = ng
	}
}

func (c *Cluster) Evoke(main model.Quale, associations map[string]model.Quale) model.Quale {
	strongestQuale := model.NewQuale()
	for name, group := range c.groups {
		if association, ok := associations[name]; ok {
			q := group.Evoke(main, association)
			if q.Strength() > strongestQuale.Strength() {
				strongestQuale = q
			}
		}
	}
	return strongestQuale
}
