package model

import "time"

// Dispatch is a model of the broadcast signal lines for communication from and to different modules. A dispatch
// communicates qualic information from a source module to subscribed modules. As an example, a dispatch for vision would communicate
// visual qualia to the auditory system, meanwhile an auditory dispatch communicates auditory qualia to the visual module.
// Only the Dispatch should communicate internal quale into modules, both associatively and main.
type Dispatch struct {
	currentQuale       Quale
	subscriberCallback []func(q Quale)
}

func (d *Dispatch) RegisterSubscriber(callback func(q Quale)) {
	d.subscriberCallback = append(d.subscriberCallback, callback)
}

func (d *Dispatch) Distribute(q Quale) {
	if q.Strength() > d.currentQuale.Strength() {
		for _, callback := range d.subscriberCallback {
			callback(q)
		}
		d.currentQuale = q
	}
}

func (d *Dispatch) Timeout(decayTime int) {
	for {
		d.currentQuale.Decay()
		time.Sleep(time.Second * time.Duration(decayTime))
	}
}
