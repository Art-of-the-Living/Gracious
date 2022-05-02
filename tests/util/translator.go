package util

import "github.com/Art-of-the-Living/gracious/base"

type JsonSignalData struct {
	Signals []JsonSignal `json:"signals"`
}

type JsonSignal struct {
	Name     string        `json:"name"`
	Features []JsonFeature `json:"features"`
}

func (js *JsonSignal) ToDistributedSignal() base.DistributedSignal {
	tmp := base.NewDistributedSignal(js.Name)
	for _, feature := range js.Features {
		tmp.Features[base.Address{X: feature.X, Y: feature.Y}] = feature.Value
	}
	return tmp
}

type JsonFeature struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Value int `json:"value"`
}
