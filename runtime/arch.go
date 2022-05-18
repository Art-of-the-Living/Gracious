package runtime

import (
	services2 "github.com/Art-of-the-Living/gracious/runtime/services"
	"sync"
	"time"
)

type Architecture struct {
	Running  chan bool
	services map[string]services2.Service
}

func NewArchitecture() *Architecture {
	arch := Architecture{
		services: make(map[string]services2.Service),
		Running:  make(chan bool),
	}
	return &arch
}

func (arch *Architecture) AddService(service services2.Service) {
	arch.services[service.GetName()] = service
}

func (arch *Architecture) RemoveService(service services2.Service) {
	delete(arch.services, service.GetName())
}

func (arch *Architecture) Update() {
	var wg sync.WaitGroup
	for _, srv := range arch.services {
		wg.Add(1)
		go srv.Broadcast(&wg)
	}
	wg.Wait()
	for _, srv := range arch.services {
		wg.Add(1)
		go srv.Listen(&wg)
	}
	wg.Wait()
}

func (arch *Architecture) Start(delay int) {
	state := true
	for state {
		arch.Update()
		time.Sleep(time.Millisecond * time.Duration(delay))
		select {
		case state = <-arch.Running:
		default:
		}
	}
}
