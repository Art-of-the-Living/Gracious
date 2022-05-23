package tests

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious"
	"github.com/Art-of-the-Living/gracious/io"
	"github.com/Art-of-the-Living/gracious/tests/tools"
	"testing"
)

func iterate(g gracious.Group, a gracious.QualitativeSignal, b gracious.QualitativeSignal, iterations int) gracious.QualitativeSignal {
	var evocation gracious.QualitativeSignal
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
	bg := gracious.NewBasicGroup("testGroup1")
	bg.CorrelationThreshold = 5
	colorJSA := io.JsonFromFileName("data/colorA.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := io.JsonFromFileName("data/wordA.json")
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
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("green")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("blue")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("yellow")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("cyan")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("magenta")
	iterate(bg, gracious.NewQualitativeSignal("void"), wordReader.Evoke(), testingIterations)
}

func TestAdvancedGroup(t *testing.T) {
	trainingIterations := 12 // Number of times to train on each signal
	testingIterations := 6   // Number of times to test on each signal
	ag := gracious.NewAdvancedGroup("testingGroupA")
	ag.CorrelationThreshold = 5
	ag.GrdCorrelationThreshold = 3
	colorJSA := io.JsonFromFileName("data/colorB.json")
	colorReader := tools.NewJsonReader(colorJSA, "blue")
	wordJSA := io.JsonFromFileName("data/wordA.json")
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
	iterate(ag, gracious.NewQualitativeSignal("redTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("green")
	iterate(ag, gracious.NewQualitativeSignal("greenTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("blue")
	iterate(ag, gracious.NewQualitativeSignal("blueTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("yellow")
	iterate(ag, gracious.NewQualitativeSignal("yellowTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("cyan")
	iterate(ag, gracious.NewQualitativeSignal("cyanTest"), wordReader.Evoke(), testingIterations)
	wordReader.SetTargetSignal("magenta")
	iterate(ag, gracious.NewQualitativeSignal("magentaTest"), wordReader.Evoke(), testingIterations)
}
