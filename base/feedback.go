package base

import (
	"github.com/KennethGrace/gracious/model"
)

type Feedback struct {
	strongestQuale model.Quale
	*Cluster
}

func NewFeedback(binding string) *Feedback {
	f := Feedback{}
	f.Cluster = NewCluster(binding)
	f.strongestQuale = model.NewQuale()
	f.SetPassThrough(true)
	return &f
}

func (f *Feedback) Evoke(main model.Quale) model.Quale {
	for _, group := range f.groups {
		group.PassThrough = true
	}
	q := f.Cluster.Evoke(main)
	return q
}
