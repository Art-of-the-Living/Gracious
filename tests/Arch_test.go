package tests

import (
	"github.com/Art-of-the-Living/gracious/learners"
	"github.com/Art-of-the-Living/gracious/runtime"
	"github.com/Art-of-the-Living/gracious/tests/tools"
	"github.com/Art-of-the-Living/gracious/util"
	"testing"
	"time"
)

func TestArchBasic(t *testing.T) {
	arch := runtime.NewArchitecture()
	colorJSA := util.JsonFromFileName("data/colorA.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := util.JsonFromFileName("data/wordA.json")
	wordReader := tools.NewJsonReader(wordJSA, "blue")
	FdColor := learners.NewAdvancedGroup("FeedbackColor")
	FdColor.PassThrough = true
	FdWord := learners.NewAdvancedGroup("FeedbackWords")
	FdWord.PassThrough = true

	FdClrSrv := runtime.NewLearnerService(FdColor.Id, FdColor)
	FdWrdSrv := runtime.NewLearnerService(FdWord.Id, FdWord)
	SnsrColorSrv := runtime.NewSensorService(colorReader.Id, colorReader)
	SnsrWordSrv := runtime.NewSensorService(wordReader.Id, wordReader)
	SnsrColorSrv.AddListener(FdClrSrv.Main)
	SnsrWordSrv.AddListener(FdWrdSrv.Main)
	arch.AddService(FdClrSrv)
	arch.AddService(FdWrdSrv)
	arch.AddService(SnsrWordSrv)
	arch.AddService(SnsrColorSrv)
	go arch.Start(500)
	time.Sleep(time.Second * 5)
	arch.Running = false
}
