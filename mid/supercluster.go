package mid

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/base"
)

// A SuperCluster is a set of clusters, such that each cluster receives the same associational inputs, but
// different main inputs. A SuperCluster is meant, primarily, for processing a slice of DistributedSignal values
// in parallel. Once learning is complete the SuperCluster can reproduce the DistributedSignal slice by associational
// evocations, even when only part or none of the main sequence is present.
type SuperCluster struct {
	Id          string     // A unique identifier for this SuperCluster
	clusters    []*Cluster // A slice of Cluster instances for evaluating slices of DistributedSignal values.
	PassThrough bool       // Determines whether the main signals should be passed through to the output
	WTA         bool       // Determines whether the result of each cluster should undergo a Winner Takes All
}

func NewSuperCluster(id string) *SuperCluster {
	sc := SuperCluster{Id: id + "-SC", clusters: make([]*Cluster, 0)}
	return &sc
}

// Evoke will test the clusters groups for associational evocations for each main signal received.
func (sc *SuperCluster) Evoke(mSignals []base.DistributedSignal, aSignals ...base.DistributedSignal) []base.DistributedSignal {
	// Test the incoming signal for building new clusters
	additionalClusterCount := len(mSignals) - len(sc.clusters)
	for i := 0; i < additionalClusterCount; i++ {
		pos := len(sc.clusters)
		nc := NewCluster(fmt.Sprint(sc.Id, "#", pos))
		nc.PassThrough = sc.PassThrough
		sc.clusters = append(sc.clusters, nc)
	}
	// Create a fresh slice of DistributedSignals and test for pass through
	nds := make([]base.DistributedSignal, 0, len(sc.clusters))
	if sc.PassThrough {
		nds = append(nds, mSignals...)
	}
	for i, cluster := range sc.clusters {
		if i < len(mSignals) {
			if i < len(nds) {
				nds[i] = cluster.Evoke(mSignals[i], aSignals...)
			} else {
				nds = append(nds, cluster.Evoke(mSignals[i], aSignals...))
			}
		} else {
			nds = append(nds, cluster.Evoke(base.NewDistributedSignal("void"), aSignals...))
		}
		if sc.WTA {
			nds[i].WinnersTakeAll(0)
		}
	}
	return nds
}
