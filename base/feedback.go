package base

import (
	"github.com/KennethGrace/gracious/model"
)

type Feedback struct {
	*Cluster
}

func NewFeedback(binding string) *Feedback {
	f := Feedback{}
	f.Cluster = NewCluster(binding)
	f.SetPassThrough(true)
	return &f
}

func (f *Feedback) Evoke(main model.Quale) model.Quale {
	strongestQuale := model.NewQuale()
	for _, group := range f.groups {
		group.PassThrough = true
		q := group.Evoke(main)
		if q.Strength() > strongestQuale.Strength() {
			strongestQuale = q
		}
	}
	return strongestQuale
}
