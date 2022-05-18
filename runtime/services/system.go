package services

import (
	"github.com/Art-of-the-Living/gracious/util"
	"sync"
)

// The Service is the foundation of the Gracious runtime. A service works in two
// phases; Advertisement and Analysis. First the service will advertise it's
// current QualitativeSignals to each Listener, then the service will analyze
// what has been received from other Service instances. When the advertisement
// phase occurs again the result of the analysis will be advertised.
// This can be understood simply as a "write" phase and a "read" phase.
type Service interface {
	GetName() string              // Returns a name identifying this service
	Broadcast(wg *sync.WaitGroup) // Writes service signal to each Listener
	Listen(wg *sync.WaitGroup)    // Reads from Listeners and determines service signal
}

// A Listener is a service component that accepts signals from one and ONLY ONE
// Broadcaster. In order to be accepted a Broadcaster must register to the Listener.
// Any previously bound broadcaster will be disconnected and incoming advertisements
// from them will be discarded.
type Listener struct {
	broadcaster *Broadcaster
	heldPattern util.QualitativeSignal
}

func NewListener() *Listener {
	b := Listener{}
	b.heldPattern = util.NewQualitativeSignal("listener-void")
	return &b
}
func (l *Listener) Register(from *Broadcaster) {
	if l.broadcaster != nil {
		l.broadcaster.RemoveListener(l)
	}
	l.broadcaster = from
}
func (l *Listener) Send(from *Broadcaster, signal util.QualitativeSignal) {
	if from == l.broadcaster {
		l.heldPattern = signal
	}
}
func (l *Listener) Read() util.QualitativeSignal {
	return l.heldPattern
}

type Broadcaster struct {
	destinations map[*Listener]*Listener // A self indexed map of Listeners
}

func NewBroadcaster() *Broadcaster {
	b := Broadcaster{destinations: make(map[*Listener]*Listener)}
	return &b
}
func (b *Broadcaster) Advertise(signal util.QualitativeSignal) {
	for _, dst := range b.destinations {
		dst.Send(b, signal)
	}
}
func (b *Broadcaster) AddListener(listener *Listener) {
	listener.Register(b)
	b.destinations[listener] = listener
}
func (b *Broadcaster) RemoveListener(listener *Listener) {
	if _, ok := b.destinations[listener]; ok {
		delete(b.destinations, listener)
	}
}
