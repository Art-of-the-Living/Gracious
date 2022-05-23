package runtime

import (
	"errors"
	"github.com/Art-of-the-Living/gracious"
	"github.com/Art-of-the-Living/gracious/io"
	"sync"
)

// The Service is the foundation of the Gracious runtime. A service works in two
// phases; Advertisement and Analysis. First the service will advertise it's
// current QualitativeSignals to each basicListener, then the service will analyze
// what has been received from other Service instances. When the advertisement
// phase occurs again the result of the analysis will be advertised.
// This can be understood simply as a "write" phase and a "read" phase.
type Service interface {
	GetListener(string) (Listener, error) // Returns a Listener from the service, if it exists
	GetName() string                      // Returns a name identifying this service
	Listen(wg *sync.WaitGroup)            // Service Listeners are called to listen
	Update(wg *sync.WaitGroup)            // Updates the Service loop
	Stop()                                // Ends the Service loop
}

type Component interface {
	GetId() string
}

// A baseService is a useful, but incomplete implementation of the Service interface.
type baseService struct {
	name          string
	active        bool
	activeChannel chan bool
	listeners     map[string]Listener
}

// newBaseService creates and returns the baseService
func newBaseService(name string) baseService {
	bs := baseService{
		name:          name,
		active:        true,
		activeChannel: make(chan bool),
		listeners:     make(map[string]Listener),
	}
	return bs
}

// GetName returns the name of the Service
func (bs *baseService) GetName() string {
	return bs.name
}

// GetListener returns the Listener of the specified name from the baseService
func (bs *baseService) GetListener(name string) (Listener, error) {
	if value, ok := bs.listeners[name]; ok {
		return value, nil
	} else {
		return nil, errors.New("call to get listener " + name + " , but listener does not exist")
	}
}

// Listen updates the Listener instances on the baseService
func (bs *baseService) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, listener := range bs.listeners {
		listener.Listen()
	}
}

// IsActive determines if the Service is still active. (Has it been stopped?)
func (bs *baseService) IsActive() bool {
	select {
	case bs.active = <-bs.activeChannel:
	default:
	}
	return bs.active
}

// Stop sets the baseService to inactive
func (bs *baseService) Stop() {
	bs.activeChannel <- false
}

type LearnerService struct {
	pattern gracious.QualitativeSignal
	group   gracious.Group
	baseService
}

func NewLearnerService(group gracious.Group) *LearnerService {
	a := LearnerService{group: group, pattern: gracious.NewQualitativeSignal(group.GetId() + "-void")}
	a.baseService = newBaseService(group.GetId())
	a.listeners["main"] = newBasicListener()
	a.listeners["associative"] = newBasicListener()
	return &a
}

func (a *LearnerService) Update(wg *sync.WaitGroup) {
	defer wg.Done()
	if a.IsActive() {
		// evoke the learning group
		a.pattern = a.group.Evoke(a.listeners["main"].GetPattern(), a.listeners["associative"].GetPattern())
	}
}

func (a *LearnerService) GetPattern() gracious.QualitativeSignal {
	return a.pattern
}

type SensorService struct {
	pattern gracious.QualitativeSignal
	sensor  io.Sensor
	baseService
}

func NewSensorService(sensor io.Sensor) *SensorService {
	s := SensorService{sensor: sensor, pattern: gracious.NewQualitativeSignal(sensor.GetId() + "void")}
	s.baseService = newBaseService(sensor.GetId())
	return &s
}

func (s *SensorService) Update(wg *sync.WaitGroup) {
	defer wg.Done()
	if s.IsActive() {
		s.pattern = s.sensor.Evoke()
	}
}

func (s *SensorService) GetPattern() gracious.QualitativeSignal {
	return s.pattern
}

type CompositionService struct {
	composition gracious.QualitativeSignal
	baseService
}

func NewCompositionService(name string) *CompositionService {
	c := CompositionService{composition: gracious.NewQualitativeSignal(name + "void")}
	c.baseService = newBaseService(name)
	c.listeners["a"] = newBasicListener()
	c.listeners["b"] = newBasicListener()
	return &c
}

func (c *CompositionService) Update(wg *sync.WaitGroup) {
	defer wg.Done()
	if c.IsActive() {
		//c.composition = gracious.Composite(c.listeners["a"].GetPattern(), c.listeners["b"].GetPattern())
	}
}

func (c *CompositionService) GetPattern() gracious.QualitativeSignal {
	return c.composition
}

type AutoMemoryService struct {
	pattern gracious.QualitativeSignal
	group   gracious.Group
	baseService
}

func NewAutoMemoryService(name string, group gracious.Group) *AutoMemoryService {
	a := AutoMemoryService{group: group, pattern: gracious.NewQualitativeSignal(name + "void")}
	a.baseService = newBaseService(name)
	a.listeners["main"] = newBasicListener()
	return &a
}

func (a *AutoMemoryService) Update(wg *sync.WaitGroup) {
	defer wg.Done()
	if a.IsActive() {
		main := a.listeners["main"].GetPattern()
		a.pattern = a.group.Evoke(main, main)
	}
}

func (a *AutoMemoryService) GetPattern() gracious.QualitativeSignal {
	return a.pattern
}
