package gracious

import (
	"sync"
)

// A Group is the atomic unit of learning. A Group should be capable of learning
// an association between a pattern and a main signal. A Group should be capable
// of re-emitting a main signal in the absence of the main signal provided the
// associated signal pattern. A Group should produce Match, Mismatch, and Novelty
// patterns depending on the relationship between the main signal and the associative
// signal at a moment of time, T.
type Group interface {
	GetId() string
	GetFirePattern() QualitativeSignal
	GetMatchPattern() QualitativeSignal
	GetMisMatchPattern() QualitativeSignal
	GetMatchLevel() int
	GetNoveltyPattern() QualitativeSignal
	GetNoveltyLevel() int
	Evoke(main, association QualitativeSignal) QualitativeSignal
	AsyncEvoke(main, association QualitativeSignal, wg *sync.WaitGroup) QualitativeSignal
}

// BasicGroup is a set of neurons with a specific associative QualitativeSignal
// type input and a specific main QualitativeSignal type input and output. The
// BasicGroup controls the learning threshold for the Neurons, as well as the
// pass through switch for whether the main signal should reappear in the output.
// The BasicGroup associates two qualitative signals with each other, of either
// 1->N or N->1. A Basic Group cannot associate M->N without interference. For
// M->N signal association the AdvancedGroup must be used.
type BasicGroup struct {
	id                   string              // The name of this group of Neurons
	neurons              map[Address]*neuron // The Neurons which compose this BasicGroup
	pattern              QualitativeSignal   // The active firing Pattern of this BasicGroup after evocation
	PassThrough          bool                // Determines if the main signal pattern should pass through to the output
	WTA                  int                 // Determines if the output of the group should undergo a WTA
	CorrelationThreshold int                 // Determines the threshold for synaptic learning in this group
}

// NewBasicGroup returns a new BasicGroup instance with an empty map of neuron instances
func NewBasicGroup(id string) *BasicGroup {
	neurons := make(map[Address]*neuron)
	ng := BasicGroup{
		neurons: neurons,
		id:      id,
	}
	return &ng
}

func (g *BasicGroup) GetId() string {
	return g.id
}

// GetFirePattern returns the actively evoked firing pattern of this group. Once retrieved this value is reset.
func (g *BasicGroup) GetFirePattern() QualitativeSignal {
	return g.pattern
}

// GetMatchPattern returns a QualitativeSignal where each feature indicates the match condition of a neuron in the Group
func (g *BasicGroup) GetMatchPattern() QualitativeSignal {
	matchPattern := NewQualitativeSignal(g.id + "-match")
	for addr, neuron := range g.neurons {
		if neuron.match {
			matchPattern.Features[addr] = 1
		}
	}
	return matchPattern
}

// GetMisMatchPattern returns a QualitativeSignal where each feature indicates the mismatch condition of a neuron in the Group
func (g *BasicGroup) GetMisMatchPattern() QualitativeSignal {
	matchPattern := NewQualitativeSignal(g.id + "-mismatch")
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

func (g *BasicGroup) GetNoveltyPattern() QualitativeSignal {
	noveltyPattern := NewQualitativeSignal(g.id + "-novelty")
	for addr, neuron := range g.neurons {
		if neuron.novelty {
			noveltyPattern.Features[addr] = 1
		}
	}
	return noveltyPattern
}

func (g *BasicGroup) GetNoveltyLevel() int {
	noveltyLevel := 0
	for _, neuron := range g.neurons {
		if neuron.novelty {
			noveltyLevel++
		} else {
			noveltyLevel--
		}
	}
	return noveltyLevel
}

// Evoke will test the BasicGroup for an associative evocation pattern. During
// evocation neuron instances will be grown and trained. The evocation will cause
// an update to the match signals which are retrievable with GetMatchPattern,
// GetMisMatchPattern, and GetMatchLevel. The pattern is returned, but can also
// be retrieved via GetPattern.
func (g *BasicGroup) Evoke(main, association QualitativeSignal) QualitativeSignal {
	if g.PassThrough {
		g.pattern = main
	} else {
		g.pattern = NewQualitativeSignal(main.Id)
	}
	// Test the incoming signal for building new neurons
	for addr := range main.Features {
		if _, ok := g.neurons[addr]; !ok { // Does the BasicGroup have a neuron at the main feature address
			g.neurons[addr] = newNeuron() // If not, create a new neuron
		}
	}
	// Test each neuron for firing strength.
	var wg sync.WaitGroup
	for addr, neuron := range g.neurons {
		wg.Add(1)
		go neuron.asyncEvoke(main.Features[addr], association, g.CorrelationThreshold, &wg)
	}
	wg.Wait()
	// Retrieve the firing strength of each neuron and adjust the firing Pattern accordingly
	for address, neuron := range g.neurons {
		if neuron.axon > 0 {
			g.pattern.Features[address] += neuron.axon
		}
	}
	if g.WTA >= 0 {
		g.pattern.WinnerTakesAll(g.WTA)
	}
	return g.pattern
}

// AsyncEvoke will Evoke this Group as a member of a WaitGroup
func (g *BasicGroup) AsyncEvoke(main, association QualitativeSignal, wg *sync.WaitGroup) QualitativeSignal {
	defer wg.Done()
	return g.Evoke(main, association)
}

// An AdvancedGroup handles more complex relationships than a BasicGroup. While a
// BasicGroup is quick to learn and less memory intensive than an AdvancedGroup,
// a BasicGroup can not resolve N->M mappings. An AdvancedGroup uses a very large
// set of extra neurons to map from N->1 and then an internal BasicGroup handles
// the 1->M.
//
// The properties of the N->1 neuron set, named the grandmother set, rely on the
// mutual exclusion of patterns. There must be one neuron for every possible
// combination of association signals, and only one neuron will identify each of
// those possible patterns for evocation in the second group.
type AdvancedGroup struct {
	grdNeurons              []*neuron // The neuron instances that compose this AdvancedGroup's N->1 map
	GrdCorrelationThreshold int       // Determines the threshold for synaptic learning in the grandmother set
	*BasicGroup                       // The component BasicGroup
}

// NewAdvancedGroup returns a new AdvancedGroup instance with an empty map of grandmother neurons
func NewAdvancedGroup(id string) *AdvancedGroup {
	g := AdvancedGroup{grdNeurons: []*neuron{newNeuron()}}
	g.BasicGroup = NewBasicGroup(id)
	return &g
}

// Evoke will test the AdvancedGroup for an associative evocation pattern. The
// Advanced Group will grow an additional neuron for each additional possible
// signal pattern. Therefore, a neuron must only be grown if there is no evocation by the existing set and
// when the previously grown neuron is done learning.
func (g *AdvancedGroup) Evoke(main QualitativeSignal, association QualitativeSignal) QualitativeSignal {
	grandmotherSignal := NewQualitativeSignal(g.id + "-grandmother")
	var wg sync.WaitGroup
	for _, neuron := range g.grdNeurons {
		wg.Add(1)
		go neuron.asyncEvoke(1, association, g.GrdCorrelationThreshold, &wg)
	}
	wg.Wait()
	// Retrieve the firing strength of each neuron and adjust the firing Pattern accordingly
	for i, neuron := range g.grdNeurons {
		if neuron.axon > 0 {
			grandmotherSignal.Features[Address{X: i}] += neuron.axon
		}
	}
	grandmotherSignal.WinnerTakesAll(0)
	// Test for new neuron growth
	topNeuron := g.grdNeurons[len(g.grdNeurons)-1]
	if topNeuron.getSumOfWeights() > 0 {
		topNeuron.learningEnabled = false
		if len(grandmotherSignal.Features) == 0 {
			g.grdNeurons = append(g.grdNeurons, newNeuron())
		}
	}
	// Send the main and grandmother signal through the basic neuron group
	g.pattern = g.BasicGroup.Evoke(main, grandmotherSignal)
	return g.pattern
}

// AsyncEvoke will Evoke this Group as a member of a WaitGroup
func (g *AdvancedGroup) AsyncEvoke(main QualitativeSignal, association QualitativeSignal, wg *sync.WaitGroup) QualitativeSignal {
	defer wg.Done()
	return g.Evoke(main, association)
}

// neuron models the unary behaviour of a single neuron. A neuron is only tangibly useful as a component part of a
// system of Neurons. The goal of each individual neuron is to form an association between synaptic inputs and the
// Neurons "firing" state. The neuron should fire, if and only if, the synaptic inputs
type neuron struct {
	// Internal Attributes
	synapses        map[Address]*Synapse
	axon            int
	match           bool
	novelty         bool
	learningEnabled bool
}

func newNeuron() *neuron {
	synapses := make(map[Address]*Synapse)
	n := neuron{synapses: synapses, learningEnabled: true}
	return &n
}

// getSumOfWeights returns the amount of synapses which have learnt their weight values. In a bipolar system, this is
// equal to the difference between the number of synapses and the true sum of weights.
func (n *neuron) getSumOfWeights() int {
	weightSum := 0
	count := len(n.synapses)
	for _, syn := range n.synapses {
		weightSum += syn.weightValue
	}
	return count + weightSum // Fancy.
}

// evoke tests the neuron for firing and writes the fired value to the 'axon'
// channel. If the firing state does not evoke in the presence of the training
// signal, the synaptic association trains itself.
func (n *neuron) evoke(training int, associative QualitativeSignal, correlation int) {
	sum := 0
	// Test the neuron synaptic associative evocations, if there is not a synapse present to handle the association
	// feature then a new synapse will be made.
	for featureAddress, feature := range associative.Features {
		if syn, ok := n.synapses[featureAddress]; ok {
			value := syn.Evoke(feature)
			sum += value
		} else {
			n.synapses[featureAddress] = NewSynapse()
		}
	}
	// Training should occur on the condition of a novelty state being produced by
	// the current system and only when learning has been enabled
	if n.learningEnabled {
		for featureAddress, feature := range associative.Features {
			if syn, ok := n.synapses[featureAddress]; ok {
				if (sum <= 0) && (training != 0) {
					n.novelty = true
					syn.Train(training, feature, correlation)
				}
			}
		}
	}
	// In the case that both signals are the same polarity, match is true.
	// In the case that both signals are of different polarity, match is false.
	n.match = ((sum > 0) && (training > 0)) || ((sum <= 0) && (training <= 0))
	n.axon = sum
}

// asyncEvoke will evoke this neuron as a member of a WaitGroup
func (n *neuron) asyncEvoke(training int, associative QualitativeSignal, correlation int, wg *sync.WaitGroup) {
	defer wg.Done()
	n.evoke(training, associative, correlation)
}

// The Synapse performs the crucial job of connecting associations to neuron groups. Each synapse has a certain weight
// which is either negative 1 or positive one. Once the weight has been set to one via +3:-1 Hebbian learning, the weight
// must be reset by a raise in the correlation sum threshold value.
type Synapse struct {
	// Internal Attributes
	correlationSum int
	weightValue    int
}

// NewSynapse initializes a new Synapse with a -1 weight value and a 0 correlation sum. A pointer to the Synapse is
// returned.
func NewSynapse() *Synapse {
	syn := &Synapse{weightValue: -1, correlationSum: 0}
	return syn
}

// Evoke is invoked at time, T, to return the result of the associational
// operation between the feature signal and the synaptic weight.
func (syn *Synapse) Evoke(association int) int {
	return association * syn.weightValue
}

// Train trains the synapse for later evocation. For learning we use an optimized
// correlative Hebbian learning algorithm for training, which prioritizes
// bit-shifting for multiplications and performs +3:-1 incremental steps.
func (syn *Synapse) Train(training int, association int, correlation int) {
	if syn.correlationSum > correlation {
		syn.weightValue = 1
	} else {
		syn.weightValue = -1
		syn.correlationSum += 4 * association * training
		syn.correlationSum -= association
	}
}
