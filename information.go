package gracious

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

// A QualitativeSignal is the smallest unit of information in Gracious. A QualitativeSignal is a form of data
// representation which does not translate signal from its original form, as is necessary in Digital Signal
// Processing. Instead, all signal manipulation in Gracious deals with the dynamic association and associational
// evocation of distributed signals with one another.
//
// A helpful example of a distributed signal is the phenomena of sight. Each sensory neuron for a given point on the
// contact surface is tuned to a different frequency of light. A distributed signal would be the exact firing Pattern
// at that point. If we say that we can determine, in a given contact point, 3 distinct color values, then a distributed
// signal of size 3. In this situation the signal could be in any of 2^n, 8, states.
type QualitativeSignal struct {
	Id       string          // A descriptive name for this signal. Useful in identification of this signal.
	Novelty  int             // The sum of all the novelty events in the production of this firing Pattern.
	MisMatch int             // The sum of all the mismatches in the production of this firing Pattern.
	Features map[Address]int // The complete set of active features in the QualitativeSignal.
}

// NewQualitativeSignal returns a new QualitativeSignal initialized with no Features set. The QualitativeSignal's size at this point is 0,
// and it occupies very little space in the system. As Features are set more, more will populate the quale. Zero values
// will never be set, as the 0 is assumed by the absence of a feature.
func NewQualitativeSignal(name string) QualitativeSignal {
	return QualitativeSignal{Id: name + "-Sig", Features: make(map[Address]int)}
}

// WinnerTakesAll forces the Features in the QualitativeSignal to fight for dominance and only the strongest features
// will remain present in the signal. The gap parameter permits a level of tolerance for features which almost meet
// with max threshold. No signal beneath 1 will ever be passed through. Signals with values above 4, will be reduced
// by half.
func (q *QualitativeSignal) WinnerTakesAll(gap int) {
	max := 1
	for _, feature := range q.Features {
		if feature > max {
			max = feature
		}
	}
	for addr, feature := range q.Features {
		if feature < (max - gap) {
			delete(q.Features, addr)
		} else {
			if feature >= 4 {
				q.Features[addr] = feature / 2
			}
		}
	}
}

// Composite will add all the features of each signal to the features of Q. Where-ever there is overlap the features are summed.
func (q *QualitativeSignal) Composite(signals ...QualitativeSignal) {
	for _, signal := range signals {
		for addr, feature := range signal.Features {
			if _, ok := q.Features[addr]; ok {
				q.Features[addr] += feature
			} else {
				q.Features[addr] = feature
			}
		}
	}
}

// Decay will reduce the strength of every signal in the QualitativeSignal by the parameter, factor. The signal level
// will continue to drop until it reaches 1. The signal level can not drop below 1, otherwise the meaning of the
// signal Pattern would be lost.
func (q *QualitativeSignal) Decay(factor int) {
	for address, feature := range q.Features {
		if feature > 1 {
			q.Features[address] -= factor
			if q.Features[address] < 1 {
				q.Features[address] = 1
			}
		}
	}
}

// Represent returns a helpful string representation of this QualitativeSignal.
func (q *QualitativeSignal) Represent() string {
	featureRepresentations := make([]string, 0)
	for address, feature := range q.Features {
		featureRepresentations = append(featureRepresentations, fmt.Sprint("<", feature, ">", address.Represent()))
	}
	if len(featureRepresentations) > 0 {
		return q.Id + "; " + strings.Join(featureRepresentations, ", ")
	} else {
		return q.Id + "; NO ACTIVITY"
	}
}
