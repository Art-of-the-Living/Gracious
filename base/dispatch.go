package base

import "github.com/KennethGrace/gracious/model"

// Dispatch is a model of the broadcast signal lines for communication from and to different modules. A dispatch
// communicates qualic information from a source module to subscribed modules. As an example, a dispatch for vision would communicate
// visual qualia to the auditory system, meanwhile an auditory dispatch communicates auditory qualia to the visual module.
// Only the Dispatch should communicate internal quale into modules, both associatively and main.
type Dispatch struct {
	name             string
	currentQuale     model.Quale
	subscribedGroups []*Group
}

func NewDispatch(name string) *Dispatch {
	return &Dispatch{name: name}
}

func (d *Dispatch) RegisterSubscriber(callback *Group) {
	d.subscribedGroups = append(d.subscribedGroups, callback)
}

func (d *Dispatch) RegisterSubscribers(callbacks []*Group) {
	for _, callback := range callbacks {
		d.subscribedGroups = append(d.subscribedGroups, callback)
	}
}

func (d *Dispatch) Distribute(q model.Quale) {
	if q.Strength() > d.currentQuale.Strength() {
		for _, callback := range d.subscribedGroups {
			callback.SetAssociation(q)
		}
		d.currentQuale = q
	}
	d.currentQuale.Decay()
}
