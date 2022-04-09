package model

// Broadcaster is a model of the broadcast signal lines for communication from and to different modules. A broadcast
// communicates qualic information, but most importantly a broadcast may only hold one qualic signal at a time. Entry
// into the broadcast is limited across time, T, by a winner-takes-all circuit.
type Broadcaster struct {
	broadcasts []Quale
}
