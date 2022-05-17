package learners

import (
	"github.com/Art-of-the-Living/gracious/util"
	"sync"
)

type Group interface {
	GetPattern() util.QualitativeSignal
	GetMatchPattern() util.QualitativeSignal
	GetMisMatchPattern() util.QualitativeSignal
	GetMatchLevel() int
	Evoke(main util.QualitativeSignal, association util.QualitativeSignal) util.QualitativeSignal
	AsyncEvoke(main util.QualitativeSignal, association util.QualitativeSignal, wg *sync.WaitGroup) util.QualitativeSignal
}

// BasicGroup is a set of neurons with a specific associative QualitativeSignal
// type input and a specific main QualitativeSignal type input and output. The
// BasicGroup controls the learning threshold for the Neurons, as well as the
// pass through switch for whether the main signal should reappear in the output.
// The BasicGroup associates two qualitative signals with each other, of either
// 1->N or N->1. A Basic Group cannot associate M->N without interference. For
// M->N signal association the AdvancedGroup must be used.
type BasicGroup struct {
	Id                   string                   // The name of this group of Neurons
	neurons              map[util.Address]*Neuron // The Neurons which compose this BasicGroup
	pattern              util.QualitativeSignal   // The active firing Pattern of this BasicGroup after evocation
	PassThrough          bool                     // Determines if the main signal pattern should pass through to the output
	CorrelationThreshold int                      // Determines the threshold for synaptic learning in this group
}

// NewBasicGroup returns a new BasicGroup instance with an empty map of Neuron instances
func NewBasicGroup(id string) *BasicGroup {
	neurons := make(map[util.Address]*Neuron)
	ng := BasicGroup{
		neurons: neurons,
		Id:      id,
	}
	return &ng
}

// GetPattern returns the actively evoked firing pattern of this group
func (g *BasicGroup) GetPattern() util.QualitativeSignal {
	return g.pattern
}

// GetMatchPattern returns a QualitativeSignal where each feature indicates the match condition of a Neuron in the Group
func (g *BasicGroup) GetMatchPattern() util.QualitativeSignal {
	matchPattern := util.NewQualitativeSignal(g.Id + "-match")
	for addr, neuron := range g.neurons {
		if neuron.match {
			matchPattern.Features[addr] = 1
		}
	}
	return matchPattern
}

// GetMisMatchPattern returns a QualitativeSignal where each feature indicates the mismatch condition of a Neuron in the Group
func (g *BasicGroup) GetMisMatchPattern() util.QualitativeSignal {
	matchPattern := util.NewQualitativeSignal(g.Id + "-mismatch")
	for addr, neuron := range g.neurons {
		if !neuron.match {
			matchPattern.Features[addr] = 1
		}
	}
	return matchPattern
}

// GetMatchLevel returns the number of matches which occurred during the latest call to Evoke this Group.
// A negative match level indicates mismatches.
func (g *BasicGroup) GetMatchLevel() int {
	matchLevel := 0
	for _, neuron := range g.neurons {
		if neuron.match {
			matchLevel++
		} else {
			matchLevel--
		}
	}
	return matchLevel
}

// Evoke will test the BasicGroup for an associative evocation pattern. During
// evocation Neuron instances will be grown and trained. The evocation will cause
// an update to the match signals which are retrievable with GetMatchPattern,
// GetMisMatchPattern, and GetMatchLevel. The pattern is returned, but can also
// be retrieved via GetPattern.
func (g *BasicGroup) Evoke(main util.QualitativeSignal, association util.QualitativeSignal) util.QualitativeSignal {
	if g.PassThrough {
		g.pattern = main
	} else {
		g.pattern = util.NewQualitativeSignal(main.Id)
	}
	// Test the incoming signal for building new neurons
	for addr := range main.Features {
		if _, ok := g.neurons[addr]; !ok { // Does the BasicGroup have a Neuron at the main feature address
			g.neurons[addr] = NewNeuron() // If not, create a new neuron
		}
	}
	// Test each neuron for firing strength.
	var wg sync.WaitGroup
	for addr, neuron := range g.neurons {
		wg.Add(1)
		go neuron.AsyncEvoke(main.Features[addr], association, g.CorrelationThreshold, &wg)
	}
	wg.Wait()
	// Retrieve the firing strength of each neuron and adjust the firing Pattern accordingly
	for address, neuron := range g.neurons {
		g.pattern.Features[address] += neuron.axon
	}
	return g.pattern
}

// AsyncEvoke will Evoke this Group as a member of a WaitGroup
func (g *BasicGroup) AsyncEvoke(main util.QualitativeSignal, association util.QualitativeSignal, wg *sync.WaitGroup) util.QualitativeSignal {
	defer wg.Done()
	return g.Evoke(main, association)
}

// An AdvancedGroup handles more complex relationships than a BasicGroup. While a
// BasicGroup is quick to learn and less memory intensive than an AdvancedGroup,
// a BasicGroup can not resolve N->M mappings. An AdvancedGroup uses a very large
// set of extra neurons to map from N->1 and then an internal BasicGroup handles
// the 1->M.
//
// The properties of the N->1 Neuron set, named the grandmother set, rely on the
// mutual exclusion of patterns. There must be one Neuron for every possible
// combination of association signals, and only one neuron will identify each of
// those possible patterns for evocation in the second group.
type AdvancedGroup struct {
	grdNeurons              []*Neuron // The Neuron instances that compose this AdvancedGroup's N->1 map
	GrdCorrelationThreshold int       // Determines the threshold for synaptic learning in the grandmother set
	*BasicGroup                       // The component BasicGroup
}

// NewAdvancedGroup returns a new AdvancedGroup instance with an empty map of grandmother neurons
func NewAdvancedGroup(id string) *AdvancedGroup {
	g := AdvancedGroup{grdNeurons: []*Neuron{NewNeuron()}}
	g.BasicGroup = NewBasicGroup(id)
	return &g
}

// Evoke will test the AdvancedGroup for an associative evocation pattern. The
// Advanced Group will grow an additional Neuron for each additional possible
// signal pattern. Therefore, a neuron must only be grown if there is no evocation by the existing set and
// when the previously grown neuron is done learning.
func (g *AdvancedGroup) Evoke(main util.QualitativeSignal, association util.QualitativeSignal) util.QualitativeSignal {
	grandmotherSignal := util.NewQualitativeSignal(g.Id + "-grandmother")
	var wg sync.WaitGroup
	for _, neuron := range g.grdNeurons {
		wg.Add(1)
		go neuron.AsyncEvoke(1, association, g.GrdCorrelationThreshold, &wg)
	}
	wg.Wait()
	// Retrieve the firing strength of each neuron and adjust the firing Pattern accordingly
	for i, neuron := range g.grdNeurons {
		grandmotherSignal.Features[util.Address{X: i}] += neuron.axon
	}
	grandmotherSignal.WinnerTakesAll(0)
	// Test for new neuron growth
	topNeuron := g.grdNeurons[len(g.grdNeurons)-1]
	if topNeuron.GetSumOfWeights() > 0 {
		topNeuron.learningEnabled = false
		if len(grandmotherSignal.Features) == 0 {
			g.grdNeurons = append(g.grdNeurons, NewNeuron())
		}
	}
	g.pattern = g.BasicGroup.Evoke(main, grandmotherSignal)
	return g.pattern
}

// AsyncEvoke will Evoke this Group as a member of a WaitGroup
func (g *AdvancedGroup) AsyncEvoke(main util.QualitativeSignal, association util.QualitativeSignal, wg *sync.WaitGroup) util.QualitativeSignal {
	defer wg.Done()
	return g.Evoke(main, association)
}
