package base

import (
	"fmt"
	"strings"
)

// An Address is the identifying tag for a location of an object in the neural geometry. It is most importantly
// used in the mapping the firing patterns, but also has many other applications.
type Address struct {
	X int // The x-position on the neural architectures coordinate plane
	Y int // The y-position on the neural architectures coordinate plane
}

// Represent returns a helpful string representation of this Address.
func (a Address) Represent() string {
	return fmt.Sprint("@(", a.X, ",", a.Y, ")")
}

// A DistributedSignal is the smallest unit of information in Gracious. A DistributedSignal is a form of data
// representation which does not translate signal from its original form, as is necessary in Digital Signal
// Processing. Instead, all signal manipulation in Gracious deals with the dynamic association and associational
// evocation of distributed signals with one another.
//
// A helpful example of a distributed signal is the phenomena of sight. Each sensory neuron for a given point on the
// contact surface is tuned to a different frequency of light. A distributed signal would be the exact firing pattern
// at that point. If we say that we can determine, in a given contact point, 3 distinct color values, then a distributed
// signal of size 3. In this situation the signal could be in any of 2^n, 8, states.
type DistributedSignal struct {
	name     string          // A descriptive name for this signal. Useful in debugging and logging.
	Novelty  int             // The sum of all the novelty events in the production of this firing pattern.
	MisMatch int             // The sum of all the mismatches in the production of this firing pattern.
	Features map[Address]int // The complete set of active features in the DistributedSignal.
}

// NewDistributedSignal returns a new DistributedSignal initialized with no Features set. The DistributedSignal's size at this point is 0,
// and it occupies very little space in the system. As Features are set more, more will populate the quale. Zero values
// will never be set, as the 0 is assumed by the absence of a feature.
func NewDistributedSignal(name string) DistributedSignal {
	return DistributedSignal{name: name, Features: make(map[Address]int)}
}

// Composite will add all the features of B to the features of A. Where-ever there is overlap the features are summed.
func (q *DistributedSignal) Composite(signal DistributedSignal) {
	for addr, feature := range signal.Features {
		if _, ok := q.Features[addr]; ok {
			q.Features[addr] += feature
		} else {
			q.Features[addr] = feature
		}
	}
}

// Count returns the total number of presently active Features in the DistributedSignal.
func (q *DistributedSignal) Count() int {
	return len(q.Features)
}

// ShiftX will shift all Features along the x-axis by a specified step integer.
func (q *DistributedSignal) ShiftX(step int) {
	shifted := make(map[Address]int, len(q.Features))
	i := 0
	for addr, feature := range q.Features {
		newAddr := addr
		newAddr.X += step
		shifted[newAddr] = feature
		i++
	}
	q.Features = shifted
}

// WinnersTakeAll forces the Features in the DistributedSignal to fight for dominance and only the strongest features
// will remain present in the signal. The gap parameter permits a level of tolerance for features which almost meet
// with max threshold. No signal beneath 1 will ever be passed through.
func (q *DistributedSignal) WinnersTakeAll(gap int) {
	max := 1
	for _, feature := range q.Features {
		if feature > max {
			max = feature
		}
	}
	for addr, feature := range q.Features {
		if feature < (max - gap) {
			delete(q.Features, addr)
		}
	}
}

// Decay will reduce the strength of every signal in the DistributedSignal by the parameter, factor. The signal level
// will continue to drop until it reaches 1. The signal level can not drop below 1, otherwise the meaning of the
// signal pattern would be lost.
func (q *DistributedSignal) Decay(factor int) {
	for address, feature := range q.Features {
		if feature > 1 {
			q.Features[address] -= factor
			if q.Features[address] < 1 {
				q.Features[address] = 1
			}
		}
	}
}

// Represent returns a helpful string representation of this DistributedSignal.
func (q *DistributedSignal) Represent() string {
	featureRepresentations := make([]string, 0)
	for address, feature := range q.Features {
		featureRepresentations = append(featureRepresentations, fmt.Sprint("<", feature, ">", address.Represent()))
	}
	if len(featureRepresentations) > 0 {
		return q.name + "; " + strings.Join(featureRepresentations, ", ")
	} else {
		return q.name + "; NO ACTIVITY"
	}
}
