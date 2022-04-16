package model

// Dispatch is a model of the broadcast signal lines for communication from and to different modules. A dispatch
// communicates qualic information from a source module to subscribed modules. As an example, a dispatch for vision would communicate
// visual qualia to the auditory system, meanwhile an auditory dispatch communicates auditory qualia to the visual module.
// Only the Dispatch should communicate internal quale into modules, both associatively and main.
type Dispatch struct {
	name                string
	currentQuale        Quale
	subscriberCallbacks []func(q Quale)
}

func NewDispatch(name string) *Dispatch {
	return &Dispatch{name: name}
}

func (d *Dispatch) RegisterSubscriber(callback func(q Quale)) {
	d.subscriberCallbacks = append(d.subscriberCallbacks, callback)
}

func (d *Dispatch) RegisterSubscribers(callbacks []func(q Quale)) {
	for _, callback := range callbacks {
		d.subscriberCallbacks = append(d.subscriberCallbacks, callback)
	}
}

func (d *Dispatch) Distribute(q Quale) {
	if q.Strength() > d.currentQuale.Strength() {
		for _, callback := range d.subscriberCallbacks {
			callback(q)
		}
		d.currentQuale = q
	}
	d.currentQuale.Decay()
}
