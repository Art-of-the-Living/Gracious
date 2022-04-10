package model

import (
	"errors"
	"fmt"
)

type Quale struct {
	features []int
}

func NewQuale(size int) Quale {
	return Quale{features: make([]int, size)}
}

func (q *Quale) Size() int {
	return len(q.features)
}

// SetFeature sets the feature value at a given line of the Quale. If there is no feature available at a given line,
// then SetFeature returns an error.
func (q *Quale) SetFeature(line int, strength int) error {
	if line < len(q.features) {
		q.features[line] = strength
		return nil
	} else {
		return errors.New(fmt.Sprint("Quale is of length", len(q.features), "but", line, "was write accessed"))
	}
}

func (q *Quale) SetFeatures(array []int) {
	for i, value := range array {
		_ = q.SetFeature(i, value)
	}
}

func (q *Quale) SetQuale(instantaneousQ Quale) {
	q.SetFeatures(instantaneousQ.GetFeatures())
}

// GetFeature returns the feature value at a given line of the Quale. If there is no feature available at a given line,
// then GetFeature returns 0. This makes a quale compatible with objects larger than it's size.
func (q *Quale) GetFeature(line int) (int, error) {
	if line < len(q.features) {
		return q.features[line], nil
	} else {
		return 0, errors.New(fmt.Sprint("Quale is of length", len(q.features), "but", line, "was read accessed"))
	}
}

// GetFeatures returns the complete array of features. This function should be used sparingly as it does not test
// for length inequalities and trouble can arise when manually mapping the features array to signal lines.
func (q *Quale) GetFeatures() []int {
	return q.features
}
