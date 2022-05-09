package util

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/base"
	"github.com/Art-of-the-Living/gracious/util"
	"strings"
)

// A TextReader is a type of FunctionalSensor where each call to evoke produces a new DistributedSignal
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
	tr.SetProcessor(func() base.DistributedSignal {
		ds := base.NewDistributedSignal(fmt.Sprint(text, "#", tr.index))
		if tr.index < len(tr.text) {
			char := int(tr.text[tr.index])
			if char >= 65 && char <= 90 {
				ds.Features[base.Address{X: 0, Y: int(char) - 65}] = 1
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
