package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Translation offers an interface for creating structures capable of conversion into QualitativeSignal values.
// By default, Json encoding is included via JsonSignalArray and JsonSignal.
type Translation interface {
	ToDistributedSignals() []QualitativeSignal
}

// JsonSignalArray encodes for a sequence (or set) of JsonSignal values. A
// JsonSignalArray can be directly converted to a slice of base.QualitativeSignal
// values via ToDistributedSignals. A JsonSignalArray can also be constructed from
// a slice of QualitativeSignal values. This allows for communication of
// sequences across the network via Json encoding.
type JsonSignalArray struct {
	Id      string       `json:"id"`
	Signals []JsonSignal `json:"signals"`
}

// GetJsonSignalById returns a JsonSignal from the array that matches the id. If there is no JsonSignal
// that matches, an empty JsonSignal will be returned.
func (jsa *JsonSignalArray) GetJsonSignalById(id string) JsonSignal {
	js := JsonFromDistributedSignal(NewQualitativeSignal("void"))
	for _, signal := range jsa.Signals {
		if signal.Id == id {
			return signal
		}
	}
	return js
}

// ToDistributedSignals converts a JsonSignalArray into a slice of base.QualitativeSignal values
func (jsa *JsonSignalArray) ToDistributedSignals() []QualitativeSignal {
	var tmp = make([]QualitativeSignal, len(jsa.Signals))
	for i, signal := range jsa.Signals {
		tmp[i] = signal.ToDistributedSignal()
	}
	return tmp
}

// JsonFromDistributedSignals takes a slice of base.QualitativeSignal values and
// formats them into JsonSignalArray.
func JsonFromDistributedSignals(signals []QualitativeSignal) JsonSignalArray {
	jsa := JsonSignalArray{Signals: make([]JsonSignal, len(signals))}
	for i, signal := range signals {
		jsa.Signals[i] = JsonFromDistributedSignal(signal)
	}
	return jsa
}

func JsonFromFileName(filename string) JsonSignalArray {
	dataFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func(dataFile *os.File) {
		err := dataFile.Close()
		if err != nil {

		}
	}(dataFile)
	bytes, _ := ioutil.ReadAll(dataFile)
	var jsa JsonSignalArray
	err = json.Unmarshal(bytes, &jsa)
	if err != nil {
		return jsa
	}
	return jsa
}

// JsonSignal represents a specific QualitativeSignal encoded in Json.
type JsonSignal struct {
	Id       string        `json:"id"`
	Features []jsonFeature `json:"features"`
}

// ToDistributedSignal converts a JsonSignal into a QualitativeSignal
func (js JsonSignal) ToDistributedSignal() QualitativeSignal {
	tmp := NewQualitativeSignal(js.Id)
	for _, feature := range js.Features {
		tmp.Features[Address{X: feature.X, Y: feature.Y}] = feature.Value
	}
	return tmp
}

// JsonFromDistributedSignal formats a base.QualitativeSignal into a JsonSignal
func JsonFromDistributedSignal(signal QualitativeSignal) JsonSignal {
	tmp := JsonSignal{Features: make([]jsonFeature, len(signal.Features))}
	i := 0
	for address, feature := range signal.Features {
		tmp.Features[i] = jsonFeature{X: address.X, Y: address.Y, Value: feature}
		i++
	}
	return tmp
}

// jsonFeature represents a specific feature of a QualitativeSignal encoded in Json.
type jsonFeature struct {
	X     int `json:"X"`
	Y     int `json:"Y"`
	Value int `json:"Value"`
}
