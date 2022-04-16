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

func (c *Cluster) AddNewGroup(binding string) func(quale model.Quale) {
	ng := NewGroup(c.binding)
	ng.LearningControlSignal = 0
	ng.CorrelationThresholdSignal = 0
	c.groups[binding] = ng
	return ng.SetAssociation
}

func (c *Cluster) AddNewGroups(bindings []string) []func(quale model.Quale) {
	funcArray := make([]func(quale model.Quale), len(bindings))
	for i, binding := range bindings {
		ng := NewGroup(c.binding)
		ng.LearningControlSignal = 0
		ng.CorrelationThresholdSignal = 0
		c.groups[binding] = ng
		funcArray[i] = ng.SetAssociation
	}
	return funcArray
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
