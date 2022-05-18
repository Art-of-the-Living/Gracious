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
	FdColor.PassThrough = false
	FdColor.CorrelationThreshold = 4
	FdColor.GrdCorrelationThreshold = 4
	FdWord := learners.NewAdvancedGroup("FeedbackWords")
	FdWord.PassThrough = false
	FdWord.GrdCorrelationThreshold = 4
	FdWord.CorrelationThreshold = 4

	FdClrSrv := runtime.NewLearnerService(FdColor.Id, FdColor)
	FdWrdSrv := runtime.NewLearnerService(FdWord.Id, FdWord)
	SensorColorService := runtime.NewSensorService(colorReader.Id, colorReader)
	SensorWordService := runtime.NewSensorService(wordReader.Id, wordReader)
	SensorColorService.AddListener(FdClrSrv.Main)
	SensorWordService.AddListener(FdWrdSrv.Main)

	arch.AddService(FdClrSrv)
	arch.AddService(FdWrdSrv)
	arch.AddService(SensorWordService)
	arch.AddService(SensorColorService)
	go arch.Start(500)
	time.Sleep(time.Second * 5)
	wordReader.SetTargetSignal("red")
	colorReader.SetTargetSignal("red")
	time.Sleep(time.Second * 5)
	wordReader.SetTargetSignal("green")
	colorReader.SetTargetSignal("green")
	time.Sleep(time.Second * 5)
	arch.Running <- false
}
