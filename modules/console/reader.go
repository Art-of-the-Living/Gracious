package console

import (
	"bufio"
	"fmt"
	"github.com/KennethGrace/gracious/base"
	"github.com/KennethGrace/gracious/model"
	"github.com/KennethGrace/gracious/modules"
	"time"
)

func ASCIIToQuale(character rune) model.Quale {
	quale := model.NewQuale()
	quale.SetFeature(model.Address{Y: int(character)}, 1)
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
	FeedbackPoint    *base.Group
	AssociationPoint *base.Cluster
	AAM              *base.AutoAssociativeMemory
	Active           bool
}

func NewReadConsole(reader *bufio.Reader) *ReadConsole {
	rc := ReadConsole{}
	rc.Phenomena = ASCIIPhenomena{Reader: reader}
	rc.FeedbackPoint = base.NewGroup("reader")
	rc.FeedbackPoint.PassThrough = true
	rc.Dispatch = base.NewDispatch("reader")
	rc.AssociationPoint = base.NewCluster("reader")
	rc.AAM = base.NewAutoAssociativeMemory("reader")
	return &rc
}

// Begin initializes the ReadConsole module and starts the looping call for inputs from the phenomena
// attribute of the ReadConsole class which, if not set, will set itself to standard in at creation.
func (rc *ReadConsole) Begin(delay int) {
	rc.Active = true
	for rc.Active {
		instantaneousQ, err := rc.Phenomena.GetQuale()
		if err != nil {
			panic("Reader crashing!")
		}
		feedbackResult := rc.FeedbackPoint.Evoke(instantaneousQ)
		rc.Dispatch.Distribute(feedbackResult)
		associatedResult := rc.AssociationPoint.Evoke(feedbackResult)
		aamResult := rc.AAM.Evoke(associatedResult)
		aamResult.WinnersTakeAll(0)
		rc.FeedbackPoint.Association = aamResult
		fmt.Println("Input Quale:", feedbackResult.Represent())
		fmt.Println("Output Quale:", aamResult.Represent())
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
