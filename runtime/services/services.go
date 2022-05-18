package services

import (
	"github.com/Art-of-the-Living/gracious/learners"
	"github.com/Art-of-the-Living/gracious/util"
	"sync"
)

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

type AutoMemoryService struct {
	name  string
	Main  *Listener
	group learners.Group
	*Broadcaster
}

func NewAutoMemoryService(name string, group learners.Group) *AutoMemoryService {
	a := AutoMemoryService{name: name, group: group}
	a.Main = NewListener()
	a.Broadcaster = NewBroadcaster()
	return &a
}

func (a *AutoMemoryService) GetName() string {
	return a.name
}

func (a *AutoMemoryService) Broadcast(wg *sync.WaitGroup) {
	defer wg.Done()
	pattern := a.group.GetPattern()
	a.Advertise(pattern)
}

func (a *AutoMemoryService) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	main := a.Main.Read()
	a.group.Evoke(main, main)
}
