package base

// AutoAssociativeMemory (AAM) is used to store Quale in memory temporally. An AAM is defined by looping the main
// model.Quale back into the associative model.Quale. Therefore, a change to either the Association Quale or the
// Main Quale will result in a change to both.
type AutoAssociativeMemory struct {
	Group
}

func NewAutoAssociativeMemory(binding string) *AutoAssociativeMemory {
	aam := AutoAssociativeMemory{}
	aam.Group = *NewGroup(binding)
	return &aam
}
