package components

import "github.com/Art-of-the-Living/gracious/base"

// Psi 𝚿 is the evaluation of phenomenal observations against internal state regarding that phenomena. Psi signals are
// our mental subjective experience. Nu and Psi are often coupled together to produce the ability for the system to
// recall the Nu phenomena via an association. But because Nu perception must compete with associational evocations,
// the Psi component is a pipelined pass-through sent to a Winner-Takes-All process at the end.
type Psi struct {
	name       string
	cluster    *base.Cluster
	Main       base.DistributedSignal            // The main signal that this cluster is "observing"
	Associates map[string]base.DistributedSignal // The associative signals that compete with "observations"
}

// NewPsi creates a new Psi
func NewPsi(name string) *Psi {
	psi := Psi{name: name}
	psi.cluster = base.NewCluster(psi.GetName() + ":cluster")
	psi.cluster.PassThrough = true
	return &psi
}

// GetName returns the name of this Psi
func (psi *Psi) GetName() string {
	return "Psi:" + psi.name
}

// Evoke returns the evocation of the Psi neuron cluster given the Psi.Main and Psi.Associates DistributedSignal.
func (psi *Psi) Evoke() base.DistributedSignal {
	signal := psi.cluster.Evoke(psi.Main, psi.Associates)
	signal.WinnersTakeAll(0)
	return signal
}
