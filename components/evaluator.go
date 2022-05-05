package components

import (
	"github.com/Art-of-the-Living/gracious/base"
	"github.com/Art-of-the-Living/gracious/objects"
)

// An Evaluator produces Psi (𝚿), the evaluation of phenomenal observations against internal states regarding that phenomena.
// Evaluator signals are mentally subjective experiences. Sensors and Evaluators are often coupled together for the system's
// ability to recall the Sensor phenomena via an association. But because Sensor perception, Nu (𝛎), must compete with
// associational evocations, the Evaluator component is a pipelined pass-through sent to a Winner-Takes-All process at the end.
type Evaluator struct {
	name       string
	cluster    *objects.Cluster
	Main       base.DistributedSignal // The main signal that this cluster is "observing"
	Associates *objects.SubNet        // The associative signals that compete with "observations"
}

// NewEvaluator creates a new Evaluator
func NewEvaluator(name string) *Evaluator {
	evaluator := Evaluator{name: "Evaluator:" + name}
	evaluator.cluster = objects.NewCluster(evaluator.GetName() + ":cluster")
	evaluator.cluster.PassThrough = true
	evaluator.cluster.CorrelationThreshold = 4
	return &evaluator
}

// GetName returns the name of this Evaluator
func (e *Evaluator) GetName() string {
	return "Evaluator:" + e.name
}

func (e *Evaluator) SetDistributedSignals(signal base.DistributedSignal, net *objects.SubNet) {

}

// Evoke returns the evocation of the Evaluator neuron cluster given the Main and Associates DistributedSignal.
func (e *Evaluator) Evoke() base.DistributedSignal {
	signal := e.cluster.Evoke(e.Main, e.Associates)
	signal.WinnersTakeAll(0)
	return signal
}
