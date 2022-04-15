package util

import (
	"github.com/KennethGrace/gracious/base"
)

type Feedback struct {
	cluster *base.Cluster
}

func NewFeedback(binding string) *Feedback {
	f := Feedback{}
	f.cluster = base.NewCluster(binding)
	return &f
}
