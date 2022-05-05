package objects

// A SuperCluster is a set of clusters, such that each cluster receives the same associational inputs, but
// different main inputs. A SuperCluster is meant, primarily, for processing a slice of DistributedSignal values
// in parallel. Once learning is complete the SuperCluster can reproduce the DistributedSignal sequence by associational
// evocations, even when the full main sequence is not present.
type SuperCluster struct {
}
