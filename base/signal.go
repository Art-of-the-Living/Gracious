package base

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

// A DistributedSignal is the smallest Gracious unit of information. Once imported into the system it becomes the fundamental unit
// of subjective experience. A DistributedSignal has a dynamic structure across different Qualar types, such as vision and audition.
// At its core a DistributedSignal is just a potentially VERY large pattern of values.
type DistributedSignal struct {
	features map[Address]int
}

// NewDistributedSignal returns a new DistributedSignal initialized with no features set. The DistributedSignal's size at this point is 0,
// and it occupies very little space in the system. As features are set more, more will populate the quale. Zero values
// will never be set, as the 0 is assumed by the absence of a feature.
func NewDistributedSignal() DistributedSignal {
	return DistributedSignal{features: make(map[Address]int)}
}

// Size returns the total number of present features in the quale. It says nothing about any MAX Size.
func (q *DistributedSignal) Size() int {
	return len(q.features)
}

// SetFeature sets the feature value at a given address of the DistributedSignal. If there is no feature available at any given
// address, then SetFeature returns an error.
func (q *DistributedSignal) SetFeature(address Address, strength int) {
	if strength != 0 {
		q.features[address] = strength
	} else {
		if _, ok := q.features[address]; ok {
			delete(q.features, address)
		}
	}
}

// SetFeatures sets all the features of a quale equal to all the values at corresponding addresses.
func (q *DistributedSignal) SetFeatures(array map[Address]int) {
	for address, value := range array {
		q.SetFeature(address, value)
	}
}

// SetQuale sets the values of one quale to all the value of another without clearing the old values of the original
// DistributedSignal. The function should be accompanied by a call to DistributedSignal.Zero.
func (q *DistributedSignal) SetQuale(instantaneousQ DistributedSignal) {
	q.SetFeatures(instantaneousQ.GetFeatures())
}

// GetFeature returns the feature value at a given line of the DistributedSignal. If there is no feature available at a given line,
// then GetFeature returns 0. This makes a quale compatible with objects larger than it's size.
func (q *DistributedSignal) GetFeature(address Address) (int, error) {
	if val, ok := q.features[address]; ok {
		return val, nil
	} else {
		return 0, errors.New(fmt.Sprint("DistributedSignal is located at", &q,
			"Bad Access!", address.vectorized(), "was read accessed"))
	}
}

// GetFeatures returns the complete map of features.
func (q *DistributedSignal) GetFeatures() map[Address]int {
	return q.features
}

// AdjustFeature adjusts a feature by a given value up or down by addition or subtraction. In the event that the
// resulting value is returned to 0, the feature is removed from the DistributedSignal.
func (q *DistributedSignal) AdjustFeature(address Address, value int) {
	newFeature := q.features[address] + value
	if newFeature != 0 {
		q.features[address] = newFeature
	} else {
		delete(q.features, address)
	}
}

// ShiftX will shift all the features on the X address of the quale by the parameter "step".
func (q *DistributedSignal) ShiftX(step int) {
	oldAdders := make([]Address, len(q.features))
	shifted := make(map[Address]int, len(q.features))
	i := 0
	for addr, feature := range q.features {
		newAddr := addr
		newAddr.X += step
		shifted[newAddr] = feature
		oldAdders[i] = addr
		i++
	}
	q.SetFeatures(shifted)
	for _, addr := range oldAdders {
		delete(q.features, addr)
	}
}

// WinnersTakeAll wipes the quale of any and all features which do not meet the threshold value. By default, the
// threshold value is equal to the maximum feature. The "gap" parameter is the amount of tolerable difference between
// the maximum feature and the threshold.
func (q *DistributedSignal) WinnersTakeAll(gap int) {
	max := 0
	for _, feature := range q.features {
		if feature > max {
			max = feature
		}
	}
	for addr, feature := range q.features {
		if feature < (max - gap) {
			delete(q.features, addr)
		}
	}
}

// Clear removes every feature from the quale. Effectively resetting all values to 0 and reducing the size of the quale.
// This should be called to avoid any overwriting.
func (q *DistributedSignal) Clear() {
	for addr := range q.features {
		delete(q.features, addr)
	}
}

// Strength returns the sum of all the features present in the quale. This is most often used for importance comparisons
// to other DistributedSignal.
func (q *DistributedSignal) Strength() int {
	sum := 0
	for _, feature := range q.features {
		sum += feature
	}
	return sum
}

// Decay is essential to releasing quale from dispatch. Over time a quale which has secured the dispatch system should
// slowly be decremented until another quale takes control. High values will not lower below 1.
func (q *DistributedSignal) Decay() {
	for address, feature := range q.features {
		if feature > 1 {
			q.SetFeature(address, feature-1)
		}
	}
}

// Represent provides a method fpr printing the contents of the quale. A string value is returned which represents all
// the features in the quale at that time.
func (q *DistributedSignal) Represent() string {
	featureRepresentations := make([]string, 0)
	for address, feature := range q.features {
		featureRepresentations = append(featureRepresentations, fmt.Sprint(address.vectorized(), "(", feature, ")"))
	}
	return strings.Join(featureRepresentations, ", ")
}
