package model

import (
	"errors"
	"fmt"
	"strings"
)

// An Address is the identifying tag for a location of an object in the qualar system. Since quale have a structure
// the Address serves to add that structure.
type Address struct {
	X int
	Y int
}

// vectorized returns a slice representation of the Address
func (a Address) vectorized() []int {
	return []int{a.X, a.Y}
}

// A Quale is the smallest Gracious unit of information. Once imported into the system it becomes the fundamental unit
// of subjective experience. A Quale has a dynamic structure across different Qualar types, such as vision and audition.
// At its core a Quale is just a potentially VERY large pattern of values.
type Quale struct {
	features map[Address]int
}

// NewQuale returns a new Quale initialized with no features set. The Quale's size at this point is 0,
// and it occupies very little space in the system. As features are set more, more will populate the quale. Zero values
// will never be set, as the 0 is assumed by the absence of a feature.
func NewQuale() Quale {
	return Quale{features: make(map[Address]int)}
}

// Size returns the total number of present features in the quale. It says nothing about any MAX Size.
func (q *Quale) Size() int {
	return len(q.features)
}

// SetFeature sets the feature value at a given address of the Quale. If there is no feature available at any given
// address, then SetFeature returns an error.
func (q *Quale) SetFeature(address Address, strength int) {
	if strength != 0 {
		q.features[address] = strength
	} else {
		if _, ok := q.features[address]; ok {
			delete(q.features, address)
		}
	}
}

// SetFeatures sets all the features of a quale equal to all the values at corresponding addresses.
func (q *Quale) SetFeatures(array map[Address]int) {
	q.features = make(map[Address]int)
	for address, value := range array {
		q.SetFeature(address, value)
	}
}

// SetQuale sets the values of one quale to all the value of another without clearing the old values of the original
// Quale. The function should be accompanied by a call to Quale.Zero.
func (q *Quale) SetQuale(instantaneousQ Quale) {
	q.SetFeatures(instantaneousQ.GetFeatures())
}

// GetFeature returns the feature value at a given line of the Quale. If there is no feature available at a given line,
// then GetFeature returns 0. This makes a quale compatible with objects larger than it's size.
func (q *Quale) GetFeature(address Address) (int, error) {
	if val, ok := q.features[address]; ok {
		return val, nil
	} else {
		return 0, errors.New(fmt.Sprint("Quale is located at", &q,
			"Bad Access!", address.vectorized(), "was read accessed"))
	}
}

// GetFeatures returns the complete map of features.
func (q *Quale) GetFeatures() map[Address]int {
	return q.features
}

// AdjustFeature adjusts a feature by a given value up or down by addition or subtraction. In the event that the
// resulting value is returned to 0, the feature is removed from the Quale.
func (q *Quale) AdjustFeature(address Address, value int) {
	newFeature := q.features[address] + value
	if newFeature != 0 {
		q.features[address] = newFeature
	} else {
		delete(q.features, address)
	}
}

// Clear removes every feature from the quale. Effectively resetting all values to 0 and reducing the size of the quale.
// This should be called to avoid any overwriting.
func (q *Quale) Clear() {
	for addr := range q.features {
		delete(q.features, addr)
	}
}

// Strength returns the sum of all the features present in the quale. This is most often used for importance comparisons
// to other Quale.
func (q *Quale) Strength() int {
	sum := 0
	for _, feature := range q.features {
		sum += feature
	}
	return sum
}

// Decay is essential to releasing quale from dispatch. Over time a quale which has secured the dispatch system should
// slowly be decremented until another quale takes control. High values will not lower below 1.
func (q *Quale) Decay() {
	for address, feature := range q.features {
		if feature > 1 {
			q.SetFeature(address, feature-1)
		}
	}
}

// Represent provides a method fpr printing the contents of the quale. A string value is returned which represents all
// the features in the quale at that time.
func (q *Quale) Represent() string {
	featureRepresentations := make([]string, 0)
	for address, feature := range q.features {
		featureRepresentations = append(featureRepresentations, fmt.Sprint(address.vectorized(), "(", feature, ")"))
	}
	return strings.Join(featureRepresentations, ", ")
}
