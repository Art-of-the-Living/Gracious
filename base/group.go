package base

// Group is a set of neurons with a specific associative DistributedSignal type input and a specific main
// DistributedSignal type input and output.
type Group struct {
	Id      string                 // The name of the system this neuron group is a part of.
	neurons map[Address]*Neuron    // The Neurons which compose this Group
	Pattern chan DistributedSignal // The active firing Pattern of this Group
}

// NewGroup returns a new Group instance with an empty map of Neuron instances
func NewGroup(id string) *Group {
	neurons := make(map[Address]*Neuron)
	ng := Group{neurons: neurons,
		Id:      id + "-G",
		Pattern: make(chan DistributedSignal),
	}
	return &ng
}

// Evoke will test the neuron group for associational evocation and learning on a main signal.
// Once Evoked, the firing Pattern of the neuron group can be retrieved from GetFiringPattern.
// Evoke should only be called using `go Evoke` as Evoke will hang writing to the NeuronGroups
// Pattern.
func (g *Group) Evoke(main DistributedSignal, association DistributedSignal, cts int) {
	// Test the incoming signal for building new neurons
	for addr := range main.Features {
		if _, ok := g.neurons[addr]; !ok {
			g.neurons[addr] = NewNeuron()
		}
	}
	// Test each neuron for firing strength.
	for addr, neuron := range g.neurons {
		go neuron.Evoke(main.Features[addr], association, cts)
	}
	// Retrieve the firing strength of each neuron and adjust the firing Pattern accordingly
	for address, neuron := range g.neurons {
		sum := <-neuron.axon
		main.Features[address] += sum
	}
	g.Pattern <- main
}
