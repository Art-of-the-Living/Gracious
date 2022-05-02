package components

import "github.com/Art-of-the-Living/gracious/base"

// An Associator produces Xi (Ξ), the subconscious evaluation of all perceptual phenomena. An Associator
// associates subjective percepts, Psi (𝚿), with one another and allows each percent to be recalled by
// the presence of associated percepts. Additionally, Associators have built in auto-associative memory.
// This allows the associator to fully reproduce the signal witch may only be partly triggered by
// traditional associations.
type Associator struct {
	name               string
	associativeCluster *base.Cluster
	memoryCluster      *base.Cluster
	Main               base.DistributedSignal            // The main signal that this associator "recalls"
	Associates         map[string]base.DistributedSignal // The associative signals that recall the main signal
}

// NewAssociator creates a new Associator
func NewAssociator(name string) *Associator {
	associator := Associator{name: "Associator:" + name}
	associator.associativeCluster = base.NewCluster(associator.name + ":associative")
	associator.associativeCluster.CorrelationThreshold = 4
	associator.memoryCluster = base.NewCluster(associator.name + ":memorial")
	associator.memoryCluster.CorrelationThreshold = 3
	return &associator
}

// GetName returns the name of this Associator
func (a *Associator) GetName() string {
	return a.name
}

// Evoke returns the evocation of the Associator neuron clusters given the Main and Associates DistributedSignal.
func (a *Associator) Evoke() base.DistributedSignal {
	partial := a.associativeCluster.Evoke(a.Main, a.Associates)
	tmp := make(map[string]base.DistributedSignal)
	tmp[a.memoryCluster.GetName()+":cluster"] = partial
	full := a.memoryCluster.Evoke(partial, tmp)
	full.WinnersTakeAll(0)
	return full
}
