package util

import "github.com/Art-of-the-Living/gracious/base"

// Translation offers an interface for creating structures capable of conversion into DistributedSignal values.
// By default, Json encoding is included via JsonSignalArray and JsonSignal.
type Translation interface {
	ToDistributedSignals() []base.DistributedSignal
}

// JsonSignalArray encodes for a sequence (or set) of JsonSignal values. A
// JsonSignalArray can be directly converted to a slice of base.DistributedSignal
// values via ToDistributedSignals. A JsonSignalArray can also be constructed from
// a slice of DistributedSignal values. This allows for communication of
// sequences across the network via Json encoding.
type JsonSignalArray struct {
	Id      string       `json:"id"`
	Signals []JsonSignal `json:"signals"`
}

// ToDistributedSignals converts a JsonSignalArray into a slice of base.DistributedSignal values
func (jsa *JsonSignalArray) ToDistributedSignals() []base.DistributedSignal {
	var tmp = make([]base.DistributedSignal, len(jsa.Signals))
	for i, signal := range jsa.Signals {
		tmp[i] = signal.ToDistributedSignal()
	}
	return tmp
}

// JsonFromDistributedSignals takes a slice of base.DistributedSignal values and
// formats them into JsonSignalArray.
func JsonFromDistributedSignals(signals []base.DistributedSignal) JsonSignalArray {
	jsa := JsonSignalArray{Signals: make([]JsonSignal, len(signals))}
	for i, signal := range signals {
		jsa.Signals[i] = JsonFromDistributedSignal(signal)
	}
	return jsa
}

// JsonSignal represents a specific DistributedSignal encoded in Json.
type JsonSignal struct {
	Id       string        `json:"id"`
	Features []jsonFeature `json:"features"`
}

// ToDistributedSignal converts a JsonSignal into a DistributedSignal
func (js *JsonSignal) ToDistributedSignal() base.DistributedSignal {
	tmp := base.NewDistributedSignal(js.Id)
	for _, feature := range js.Features {
		tmp.Features[base.Address{X: feature.X, Y: feature.Y}] = feature.Value
	}
	return tmp
}

// JsonFromDistributedSignal formats a base.DistributedSignal into a JsonSignal
func JsonFromDistributedSignal(signal base.DistributedSignal) JsonSignal {
	tmp := JsonSignal{Features: make([]jsonFeature, len(signal.Features))}
	i := 0
	for address, feature := range signal.Features {
		tmp.Features[i] = jsonFeature{X: address.X, Y: address.Y, Value: feature}
		i++
	}
	return tmp
}

// jsonFeature represents a specific feature of a DistributedSignal encoded in Json.
type jsonFeature struct {
	X     int `json:"X"`
	Y     int `json:"Y"`
	Value int `json:"Value"`
}
