package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Art-of-the-Living/gracious/base"
	"github.com/Art-of-the-Living/gracious/components"
	"github.com/Art-of-the-Living/gracious/tests/util"
	"io/ioutil"
	"os"
	"testing"
)

func TestDSABasic(t *testing.T) {
	// Reading the Test Data Files
	colorDataFile, err := os.Open("data/colorA.json")
	if err != nil {
		fmt.Println(err)
	}
	defer colorDataFile.Close()
	wordDataFile, err := os.Open("data/wordA.json")
	if err != nil {
		fmt.Println(err)
	}
	defer wordDataFile.Close()
	var colorSignalData util.JsonSignalData
	var wordSignalData util.JsonSignalData
	bytes, _ := ioutil.ReadAll(colorDataFile)
	json.Unmarshal(bytes, &colorSignalData)
	bytes, _ = ioutil.ReadAll(wordDataFile)
	json.Unmarshal(bytes, &wordSignalData)

	// SetUp Neural Components
	sensorA := components.NewSensor("Color")
	sensorB := components.NewSensor("Word")
	displayCount := 8
	aSteps := 0
	aPos := 0
	sensorA.SetProcessor(func() base.DistributedSignal {
		tmp := colorSignalData.Signals[aPos].ToDistributedSignal()
		aSteps++
		if aSteps >= displayCount {
			aPos++
			aSteps = 0
			if aPos >= len(colorSignalData.Signals) {
				aPos = 0
			}
		}
		return tmp
	})
	bSteps := 0
	bPos := 0
	sensorB.SetProcessor(func() base.DistributedSignal {
		tmp := wordSignalData.Signals[bPos].ToDistributedSignal()
		bSteps++
		if bSteps >= displayCount {
			bPos++
			bSteps = 0
			if bPos >= len(wordSignalData.Signals) {
				bPos = 0
			}
		}
		return tmp
	})
	colorEvaluator := components.NewEvaluator("Color")
	wordEvaluator := components.NewEvaluator("Word")

	colorAssc := components.NewAssociator("Color")
	wordAssc := components.NewAssociator("Word")

	PsiNetwork := base.NewNetwork("Psi")
	XiNetwork := base.NewNetwork("Xi")
	count := 64
	for i := 0; i < count; i++ {
		fmt.Println("\tTHE CURRENT OBSERVATION IS", i, "OF", count)
		NuA := sensorA.Evoke()
		NuB := sensorB.Evoke()
		fmt.Println(NuA.Represent())
		fmt.Println(NuB.Represent())
		colorEvaluator.Main = NuA
		colorEvaluator.Associates = XiNetwork.GetConnections("Color")
		wordEvaluator.Main = NuB
		wordEvaluator.Associates = XiNetwork.GetConnections("Word")
		PsiColor := colorEvaluator.Evoke()
		PsiWord := wordEvaluator.Evoke()
		fmt.Println(PsiColor.Represent())
		fmt.Println(PsiWord.Represent())
		PsiNetwork.AddSignals("Color", PsiWord)
		PsiNetwork.AddSignals("Word", PsiColor)
		colorAssc.Main = PsiColor
		wordAssc.Main = PsiWord
		colorAssc.Associates = PsiNetwork.GetConnections("Color")
		wordAssc.Associates = PsiNetwork.GetConnections("Word")
		XiColor := colorAssc.Evoke()
		XiWord := wordAssc.Evoke()
		XiNetwork.AddSignals("Color", XiColor)
		XiNetwork.AddSignals("Word", XiWord)
		fmt.Println(XiColor.Represent())
		fmt.Println(XiWord.Represent())
	}
}
