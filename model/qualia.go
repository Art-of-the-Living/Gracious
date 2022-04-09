package model

type Quale struct {
	features []int
}

func NewQuale(size int) Quale {
	return Quale{features: make([]int, size)}
}

func (q *Quale) SetFeature(line int, strength int) {
	q.features[line] = strength
}
