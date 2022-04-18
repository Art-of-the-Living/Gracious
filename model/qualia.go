package model

import (
	"errors"
	"fmt"
	"strings"
)

type Address struct {
	X int
	Y int
}

func (a Address) vectorized() []int {
	return []int{a.X, a.Y}
}

type Quale struct {
	features map[Address]int
}

func NewQuale() Quale {
	return Quale{features: make(map[Address]int)}
}

func (q *Quale) Zero() {
	for addr := range q.features {
		q.features[addr] = 0
	}
}

func (q *Quale) Size() int {
	return len(q.features)
}

// SetFeature sets the feature value at a given line of the Quale. If there is no feature available at a given line,
// then SetFeature returns an error.
func (q *Quale) SetFeature(address Address, strength int) error {
	if _, ok := q.features[address]; ok {
		q.features[address] = strength
		return nil
	} else {
		q.features[address] = strength
		return errors.New(fmt.Sprint("Quale is located at", &q,
			"Bad access! ", address.vectorized(), "was write accessed"))
	}
}

func (q *Quale) SetFeatures(array map[Address]int) {
	q.features = make(map[Address]int)
	for address, value := range array {
		_ = q.SetFeature(address, value)
	}
}

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

// GetFeatures returns the complete array of features. This function should be used sparingly as it does not test
// for length inequalities and trouble can arise when manually mapping the features array to signal lines.
func (q *Quale) GetFeatures() map[Address]int {
	return q.features
}

func (q *Quale) AdjustFeature(address Address, value int) {
	q.features[address] = q.features[address] + value
}

func (q *Quale) Strength() int {
	sum := 0
	for _, feature := range q.features {
		sum += feature
	}
	return sum
}

func (q *Quale) Decay() {
	for address, feature := range q.features {
		if feature > 1 {
			_ = q.SetFeature(address, feature-1)
		}
	}
}

func (q *Quale) Represent() string {
	featureRepresentations := make([]string, 0)
	for address, feature := range q.features {
		featureRepresentations = append(featureRepresentations, fmt.Sprint(address.vectorized(), "(", feature, ")"))
	}
	return strings.Join(featureRepresentations, ", ")
}
