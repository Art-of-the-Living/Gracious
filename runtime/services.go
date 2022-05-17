package runtime

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/learners"
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
	b.heldPattern = util.NewQualitativeSignal("void")
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

type OperatorService struct {
	name     string
	Main     *Listener
	operator util.Operator
	cue      util.QualitativeSignal
	*Broadcaster
}

func NewOperatorService(operator util.Operator) *OperatorService {
	s := OperatorService{operator: operator}
	s.Main = NewListener()
	s.Broadcaster = NewBroadcaster()
	return &s
}
func (s *OperatorService) GetName() string {
	return s.name
}
func (s *OperatorService) Broadcast(wg *sync.WaitGroup) {
	defer wg.Done()
	s.Advertise(s.cue)
}
func (s *OperatorService) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	s.cue = s.operator.Execute(s.Main.Read())
}

type LearnerService struct {
	name        string
	Main        *Listener
	Associative *Listener
	group       learners.Group
	*Broadcaster
}

func NewLearnerService(name string, group learners.Group) *LearnerService {
	a := LearnerService{name: name, group: group}
	a.Main = NewListener()
	a.Associative = NewListener()
	a.Broadcaster = NewBroadcaster()
	return &a
}
func (a *LearnerService) GetName() string {
	return a.name
}
func (a *LearnerService) Broadcast(wg *sync.WaitGroup) {
	defer wg.Done()
	pattern := a.group.GetPattern()
	fmt.Println(a.GetName(), pattern.Represent())
	a.Advertise(pattern)
}
func (a *LearnerService) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	main := a.Main.Read()
	associative := a.Associative.Read()
	a.group.Evoke(main, associative)
}

type SensorService struct {
	name    string
	pattern util.QualitativeSignal
	sensor  util.Sensor
	*Broadcaster
}

func NewSensorService(name string, sensor util.Sensor) *SensorService {
	s := SensorService{sensor: sensor, name: name}
	s.pattern = util.NewQualitativeSignal(s.name + "void")
	s.Broadcaster = NewBroadcaster()
	return &s
}
func (s *SensorService) GetName() string {
	return s.name
}
func (s *SensorService) Broadcast(wg *sync.WaitGroup) {
	defer wg.Done()
	s.Advertise(s.pattern)
}
func (s *SensorService) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	s.pattern = s.sensor.Evoke()
}

type CompositionService struct {
	name        string
	A           *Listener
	B           *Listener
	composition util.QualitativeSignal
	*Broadcaster
}

func NewCompositionService(name string) *CompositionService {
	c := CompositionService{name: name}
	c.A = NewListener()
	c.B = NewListener()
	c.Broadcaster = NewBroadcaster()
	return &c
}
func (c *CompositionService) GetName() string {
	return c.name
}
func (c *CompositionService) Broadcast(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Advertise(c.composition)
}
func (c *CompositionService) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	c.composition = util.Composite(c.A.Read(), c.B.Read())
}
