package tests

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious"
	"github.com/Art-of-the-Living/gracious/tests/tools"
	"github.com/Art-of-the-Living/gracious/util"
	"testing"
)

func iterate(g gracious.Group, a gracious.QualitativeSignal, b gracious.QualitativeSignal, iterations int) gracious.QualitativeSignal {
	var evocation gracious.QualitativeSignal
	for i := 0; i < iterations; i++ {
		evocation = g.Evoke(a, b)
		evocation.WinnersTakeAll(0)
		fmt.Println(evocation.Represent())
	}
	return evocation
}

func TestBasicGroup(t *testing.T) {
	trainingIterations := 6 // Number of times to train on each signal
	testingIterations := 6  // Number of times to test on each signal
	bg := gracious.NewBasicGroup("testGroup1")
	bg.CorrelationThreshold = 5
	colorJSA := util.JsonFromFileName("data/colorA.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := util.JsonFromFileName("data/wordA.json")
	wordReader := tools.NewJsonReader(wordJSA, "blue")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("red")
	wordReader.SetTarget("red")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("cyan")
	wordReader.SetTarget("cyan")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("green")
	wordReader.SetTarget("green")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("magenta")
	wordReader.SetTarget("magenta")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("yellow")
	wordReader.SetTarget("yellow")
	iterate(bg, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	// ### END OF TRAINING ###
	wordReader.SetTarget("red")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("green")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("blue")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("yellow")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("cyan")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("magenta")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
}

func TestAdvancedGroup(t *testing.T) {
	trainingIterations := 12 // Number of times to train on each signal
	testingIterations := 6   // Number of times to test on each signal
	ag := gracious.NewAdvancedGroup("testingGroupA")
	ag.CorrelationThreshold = 5
	ag.GrdCorrelationThreshold = 3
	colorJSA := util.JsonFromFileName("data/colorB.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := util.JsonFromFileName("data/wordA.json")
	wordReader := tools.NewJsonReader(wordJSA, "blue")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("red")
	wordReader.SetTarget("red")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("cyan")
	wordReader.SetTarget("cyan")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("green")
	wordReader.SetTarget("green")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("magenta")
	wordReader.SetTarget("magenta")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	colorReader.SetTarget("yellow")
	wordReader.SetTarget("yellow")
	iterate(ag, colorReader.Evoke(), wordReader.Evoke(), trainingIterations)
	// ### END OF TRAINING ###
	wordReader.SetTarget("red")
	iterate(ag, gracious.NewQualitativeSignal("redTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("green")
	iterate(ag, gracious.NewQualitativeSignal("greenTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("blue")
	iterate(ag, gracious.NewQualitativeSignal("blueTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("yellow")
	iterate(ag, gracious.NewQualitativeSignal("yellowTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("cyan")
	iterate(ag, gracious.NewQualitativeSignal("cyanTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTarget("magenta")
	iterate(ag, gracious.NewQualitativeSignal("magentaTest"), wordReader.Evoke(), testingIterations)
}
