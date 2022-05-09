package util

import "github.com/Art-of-the-Living/gracious/base"

// A Sensor is any structure which can produce a slice of base.DistributedSignal values. The Sensor interface is
// implemented by the FunctionalSensor
type Sensor interface {
	Evoke() base.DistributedSignal
}

// A FunctionalSensor produces phenomenal experience as direct sensory
// observations. A FunctionalSensor component processes incoming raw-data into
// distributed signals that the system can begin to work with. Unlike other
// components a FunctionalSensor component doesn't have any DistributedSignal
// inputs, instead Sensors produces a DistributedSignal based on the output of an
// externally implemented function. This function can be set using SetProcessor
// and should be set before any calls to Evoke. Otherwise, Evoke will return an
// empty signal.
type FunctionalSensor struct {
	Id        string
	processor func() base.DistributedSignal
}

// NewFunctionalSensor creates a new FunctionalSensor
func NewFunctionalSensor(name string) *FunctionalSensor {
	fs := FunctionalSensor{Id: "FunctionalSensor:" + name}
	return &fs
}

// GetName returns the ID of this FunctionalSensor
func (n *FunctionalSensor) GetName() string {
	return n.Id
}

// SetProcessor sets the function that should be run when this FunctionalSensor component is evoked.
func (n *FunctionalSensor) SetProcessor(processor func() base.DistributedSignal) {
	n.processor = processor
}

// Evoke returns the result of the specified processor function. If no processor is defined, then an empty slice of
// DistributedSignal values will be returned.
func (n *FunctionalSensor) Evoke() base.DistributedSignal {
	if n.processor != nil {
		return n.processor()
	} else {
		return base.NewDistributedSignal(n.Id)
	}
}
