package tools

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/util"
	"strings"
)

// A TextReader is a type of FunctionalSensor where each call to evoke produces a new QualitativeSignal
// for each letter in the TextReader's text property. Once the string has been read out, no more signals
// will be produced. Text is considered to be only the 26 letters of the alphabet. Text is not case-sensitive.
// Text which is not a letter is interpreted to be a blank signal.
type TextReader struct {
	text  string
	index int
	util.FunctionalSensor
}

// NewTextReader creates a nex TextReader instance with the text value passed to text
func NewTextReader(text string) *TextReader {
	tr := TextReader{text: strings.ToUpper(text), index: 0}
	tr.SetProcessor(func() util.QualitativeSignal {
		ds := util.NewQualitativeSignal(fmt.Sprint(text, "#", tr.index))
		if tr.index < len(tr.text) {
			char := int(tr.text[tr.index])
			if char >= 65 && char <= 90 {
				ds.Features[util.Address{X: 0, Y: int(char) - 65}] = 1
			}
			tr.index++
		}
		return ds
	})
	return &tr
}

// Next returns true if there is any remaining text to be read
func (tr *TextReader) Next() bool {
	return tr.index < len(tr.text)
}

// Reset returns the index value of the TextReader to 0 to begin evocations again.
func (tr *TextReader) Reset() {
	tr.index = 0
}

// A JsonReader is a type of FunctionalSensor where each call to evoke produces a new QualitativeSignal
// from the set of signal data based on the currently set
type JsonReader struct {
	signals  util.JsonSignalArray
	targetId string
	util.FunctionalSensor
}

func NewJsonReader(signals util.JsonSignalArray, targetId string) *JsonReader {
	jsR := JsonReader{signals: signals, targetId: targetId}
	jsR.SetProcessor(jsR.getActiveSignal)
	jsR.Id = signals.Id
	return &jsR
}

func (jsR *JsonReader) getActiveSignal() util.QualitativeSignal {
	return jsR.signals.GetJsonSignalById(jsR.targetId).ToDistributedSignal()
}

// SetTargetSignal sets the id that should be evoked from the JsonSignalArray
func (jsR *JsonReader) SetTargetSignal(id string) {
	jsR.targetId = id
}
