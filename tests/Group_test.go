package tests

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/learners"
	"github.com/Art-of-the-Living/gracious/tests/tools"
	"github.com/Art-of-the-Living/gracious/util"
	"testing"
)

func iterate(g learners.Group, a util.QualitativeSignal, b util.QualitativeSignal, iterations int) util.QualitativeSignal {
	var evocation util.QualitativeSignal
	fmt.Println("Association: ", b.Represent())
	for i := 0; i < iterations; i++ {
		evocation = g.Evoke(a, b)
		evocation.WinnerTakesAll(0)
		fmt.Println(evocation.Represent())
	}
	return evocation
}

func TestBasicGroup(t *testing.T) {
	trainingIterations := 6 // Number of times to train on each signal
	testingIterations := 6  // Number of times to test on each signal
	bg := learners.NewBasicGroup("testGroup1")
	bg.CorrelationThreshold = 5
	colorJSA := util.JsonFromFileName("data/colorA.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := util.JsonFromFileName("data/wordA.json")
	wordReader := tools.NewJsonReader(wordJSA, "blue")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("red")
	wordReader.SetTargetSignal("red")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("cyan")
	wordReader.SetTargetSignal("cyan")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("green")
	wordReader.SetTargetSignal("green")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("magenta")
	wordReader.SetTargetSignal("magenta")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("yellow")
	wordReader.SetTargetSignal("yellow")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	// ### END OF TRAINING ###
	wordReader.SetTargetSignal("red")
	iterate(bg, util.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("green")
	iterate(bg, util.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("blue")
	iterate(bg, util.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("yellow")
	iterate(bg, util.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("cyan")
	iterate(bg, util.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("magenta")
	iterate(bg, util.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
}

func TestAdvancedGroup(t *testing.T) {
	trainingIterations := 12 // Number of times to train on each signal
	testingIterations := 6   // Number of times to test on each signal
	ag := learners.NewAdvancedGroup("testingGroupA")
	ag.CorrelationThreshold = 5
	ag.GrdCorrelationThreshold = 3
	colorJSA := util.JsonFromFileName("data/colorB.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := util.JsonFromFileName("data/wordA.json")
	wordReader := tools.NewJsonReader(wordJSA, "blue")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("red")
	wordReader.SetTargetSignal("red")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("cyan")
	wordReader.SetTargetSignal("cyan")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("green")
	wordReader.SetTargetSignal("green")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("magenta")
	wordReader.SetTargetSignal("magenta")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTargetSignal("yellow")
	wordReader.SetTargetSignal("yellow")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	// ### END OF TRAINING ###
	fmt.Println("### BEGIN TESTING ###")
	wordReader.SetTargetSignal("red")
	iterate(ag, util.NewQualitativeSignal("redTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("green")
	iterate(ag, util.NewQualitativeSignal("greenTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("blue")
	iterate(ag, util.NewQualitativeSignal("blueTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("yellow")
	iterate(ag, util.NewQualitativeSignal("yellowTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("cyan")
	iterate(ag, util.NewQualitativeSignal("cyanTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("magenta")
	iterate(ag, util.NewQualitativeSignal("magentaTest"), wordReader.Evoke(), testingIterations)
}
