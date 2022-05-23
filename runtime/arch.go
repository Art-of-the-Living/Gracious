package runtime

import (
	"sync"
	"time"
)

type Architecture struct {
	Running  chan bool
	services map[string]Service
}

func NewArchitecture() *Architecture {
	arch := Architecture{
		services: make(map[string]Service),
		Running:  make(chan bool),
	}
	return &arch
}

func (arch *Architecture) AddService(service Service) {
	arch.services[service.GetName()] = service
}

func (arch *Architecture) RemoveService(service Service) {
	delete(arch.services, service.GetName())
}

func (arch *Architecture) Update() {
	var wg sync.WaitGroup
	for _, srv := range arch.services {
		wg.Add(1)
		go srv.Update(&wg)
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
