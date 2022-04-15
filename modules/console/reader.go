package console

import (
	"bufio"
	"fmt"
	"github.com/KennethGrace/gracious/model"
	"github.com/KennethGrace/gracious/modules"
	"github.com/KennethGrace/gracious/modules/util"
	"time"
)

const ModuleName = "reader"

func ASCIIToQuale(character rune) model.Quale {
	quale := model.NewQuale()
	_ = quale.SetFeature(model.Address{Y: int(character)}, 1)
	return quale
}

type ASCIIPhenomena struct {
	Reader *bufio.Reader
}

func (p ASCIIPhenomena) GetQuale() (model.Quale, error) {
	character, _, err := p.Reader.ReadRune()
	if err != nil {
		return model.NewQuale(), err
	}
	return ASCIIToQuale(character), nil
}

// The ReadConsole grants access to a neuron group for communicating sensory information. A read console receives
// external console characters as representations of traditional ASCII characters directly into the system without
// sensory pre-processing of simulated or imported external data. This is often useful for "neural debugging" of
// the system.
type ReadConsole struct {
	modules.Module
	feedback util.Feedback
	Active   bool
}

func NewReadConsole(reader *bufio.Reader) *ReadConsole {
	p := ASCIIPhenomena{Reader: bufio.NewReader(reader)}
	rc := ReadConsole{}
	rc.Phenomena = p
	return &rc
}

// Begin initializes the ReadConsole module and starts the looping call for inputs from the phenomena
// attribute of the ReadConsole class which, if not set, will set itself to standard in at creation.
func (rc *ReadConsole) Begin() {
	q := model.NewQuale()
	rc.Active = true
	for rc.Active {
		instantaneousQ, _ := rc.Phenomena.GetQuale()
		q.SetQuale(instantaneousQ)
		fmt.Println(q)
		time.Sleep(1 * time.Second)
	}
}
