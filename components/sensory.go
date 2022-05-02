package components

import "github.com/Art-of-the-Living/gracious/base"

// A Sensor produces Nu (𝛎), phenomenal experience as direct sensory observations. A Sensor component processes incoming
// raw-data into distributed signals that the system can begin to work with. Unlike other components a Sensor component
// doesn't have any DistributedSignal inputs, instead Sensors produces a DistributedSignal based on the output of an
// externally implemented function. This function can be set using SetProcessor and should be set before any
// calls to Evoke. Otherwise, Evoke will return an empty signal.
type Sensor struct {
	name      string
	processor func() base.DistributedSignal
}

// NewSensor creates a new Sensor
func NewSensor(name string) *Sensor {
	return &Sensor{name: "Sensor:" + name}
}

// GetName returns the name of this Sensor
func (n *Sensor) GetName() string {
	return n.name
}

// SetProcessor sets the function that should be run when this Sensor component is evoked.
func (n *Sensor) SetProcessor(processor func() base.DistributedSignal) {
	n.processor = processor
}

// Evoke returns the result of the specified processor function. If no processor is defined, then an empty signal will
// be returned.
func (n *Sensor) Evoke() base.DistributedSignal {
	if n.processor != nil {
		return n.processor()
	} else {
		return base.NewDistributedSignal(n.name + ":voidPhenomena")
	}
}
