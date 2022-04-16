package console

import (
	"bufio"
	"github.com/KennethGrace/gracious/model"
	"github.com/KennethGrace/gracious/modules"
	"github.com/KennethGrace/gracious/modules/util"
	"os"
	"time"
)

func QualeToASCII(quale model.Quale) rune {
	i := 0
	for _, feature := range quale.GetFeatures() {
		if feature != 0 {
			return rune(i)
		}
		i++
	}
	return rune(i)
}

type ASCIIAction struct {
	Writer *bufio.Writer
}

func (a ASCIIAction) SetQuale(quale model.Quale) error {
	_, err := a.Writer.WriteRune(QualeToASCII(quale))
	return err
}

type WriteConsole struct {
	modules.Module
	feedback util.Feedback
	Active   bool
}

func NewWriteConsole() *WriteConsole {
	wc := WriteConsole{}
	wc.Action = ASCIIAction{Writer: bufio.NewWriter(os.Stdout)}
	return &wc
}

func (wc *WriteConsole) Begin(delay int) {
	q := model.NewQuale()
	wc.Active = true
	for wc.Active {
		_ = wc.Action.SetQuale(q)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
