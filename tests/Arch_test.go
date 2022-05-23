package tests

import (
	"github.com/Art-of-the-Living/gracious"
	"github.com/Art-of-the-Living/gracious/io"
	"github.com/Art-of-the-Living/gracious/runtime"
	"github.com/Art-of-the-Living/gracious/tests/tools"
	"testing"
	"time"
)

func TestArchBasic(t *testing.T) {
	arch := runtime.NewArchitecture()
	colorJSA := io.JsonFromFileName("data/colorB.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := io.JsonFromFileName("data/wordA.json")
	wordReader := tools.NewJsonReader(wordJSA, "blue")

	FdColor := gracious.NewAdvancedGroup("FeedbackColor")
	FdColor.PassThrough = true
	FdColor.CorrelationThreshold = 4
	FdColor.GrdCorrelationThreshold = 4
	FdWord := gracious.NewAdvancedGroup("FeedbackWords")
	FdWord.PassThrough = true
	FdWord.GrdCorrelationThreshold = 4
	FdWord.CorrelationThreshold = 4

	AssColor := gracious.NewAdvancedGroup("AssociatorColor")
	AssColor.PassThrough = false
	AssColor.CorrelationThreshold = 8
	AssColor.GrdCorrelationThreshold = 8
	AssWord := gracious.NewAdvancedGroup("AssociatorWord")
	AssWord.PassThrough = false
	AssWord.CorrelationThreshold = 8
	AssWord.GrdCorrelationThreshold = 8

	FdClrSrv := runtime.NewLearnerService(FdColor)
	FdWrdSrv := runtime.NewLearnerService(FdWord)
	AssClrSrv := runtime.NewLearnerService(AssColor)
	AssWrdSrv := runtime.NewLearnerService(AssWord)
	SensorColorService := runtime.NewSensorService(colorReader)
	SensorWordService := runtime.NewSensorService(wordReader)

	var lst runtime.Listener
	lst, _ = FdClrSrv.GetListener("main")
	lst.SetProvider(SensorColorService)
	lst, _ = FdWrdSrv.GetListener("main")
	lst.SetProvider(SensorWordService)
	lst, _ = AssClrSrv.GetListener("main")
	lst.SetProvider(FdClrSrv)
	lst, _ = AssWrdSrv.GetListener("main")
	lst.SetProvider(FdWrdSrv)

	arch.AddService(AssClrSrv)
	arch.AddService(AssWrdSrv)
	arch.AddService(FdClrSrv)
	arch.AddService(FdWrdSrv)
	arch.AddService(SensorWordService)
	arch.AddService(SensorColorService)
	runTime := 60
	go arch.Start(50)
	for i := 0; i < runTime; i++ {
		time.Sleep(time.Millisecond * 500)
		wordReader.SetTargetSignal("red")
		colorReader.SetTargetSignal("red")
		time.Sleep(time.Millisecond * 500)
		wordReader.SetTargetSignal("green")
		colorReader.SetTargetSignal("green")
		time.Sleep(time.Millisecond * 500)
		wordReader.SetTargetSignal("blue")
		colorReader.SetTargetSignal("blue")
	}
	arch.Running <- false
}
