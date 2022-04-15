package model

// Dispatch is a model of the broadcast signal lines for communication from and to different modules. A dispatch
// communicates qualic information from a source module to subscribed modules. As an example, a dispatch for vision would communicate
// visual qualia to the auditory system, meanwhile an auditory dispatch communicates auditory qualia to the visual module.
// Only the Dispatch should communicate quale into modules.
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
