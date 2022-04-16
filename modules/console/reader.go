package console

import (
	"bufio"
	"github.com/KennethGrace/gracious/base"
	"github.com/KennethGrace/gracious/model"
	"github.com/KennethGrace/gracious/modules"
	"time"
)

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
	Feedback *base.Cluster
	Active   bool
}

func NewReadConsole(reader *bufio.Reader) *ReadConsole {
	rc := ReadConsole{}
	rc.Phenomena = ASCIIPhenomena{Reader: reader}
	rc.Dispatch = model.NewDispatch("reader")
	rc.Feedback = base.NewCluster("reader")
	return &rc
}

// Begin initializes the ReadConsole module and starts the looping call for inputs from the phenomena
// attribute of the ReadConsole class which, if not set, will set itself to standard in at creation.
func (rc *ReadConsole) Begin(delay int) {
	rc.Active = true
	for rc.Active {
		instantaneousQ, _ := rc.Phenomena.GetQuale()
		rc.Dispatch.Distribute(instantaneousQ)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
