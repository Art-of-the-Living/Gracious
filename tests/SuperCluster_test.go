package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Art-of-the-Living/gracious/base"
	"github.com/Art-of-the-Living/gracious/mid"
	"github.com/Art-of-the-Living/gracious/tests/tools"
	"github.com/Art-of-the-Living/gracious/util"
	"io/ioutil"
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	var sc = mid.NewSuperCluster("test")
	sc.PassThrough = false
	sc.WTA = true
	var seq = mid.NewSequencer("test", 8)
	colorDataFile, err := os.Open("data/colorA.json")
	if err != nil {
		fmt.Println(err)
	}
	defer colorDataFile.Close()
	var colorSignalData util.JsonSignalArray
	bytes, _ := ioutil.ReadAll(colorDataFile)
	json.Unmarshal(bytes, &colorSignalData)
	var textReaders = []*tools.TextReader{
		tools.NewTextReader("Red "),
		tools.NewTextReader("Green "),
		tools.NewTextReader("Blue "),
		tools.NewTextReader("Yellow "),
		tools.NewTextReader("Cyan "),
		tools.NewTextReader("Magenta "),
	}
	i := 0
	for w, text := range textReaders {
		for d := 0; d < 3; d++ {
			for text.Next() {
				fmt.Println("\t Iteration:", i)
				value := seq.Sequence(text.Evoke())
				for _, sig := range value {
					fmt.Println("Sequence:\t", sig.Represent())
				}
				evocation := sc.Evoke(value, colorSignalData.Signals[w].ToDistributedSignal())
				for _, sig := range evocation {
					fmt.Println("SuperCluster:\t", sig.Represent())
				}
				i++
			}
			text.Reset()
		}
	}
	fmt.Println("### TESTING TRAINING ###")
	for _, signal := range colorSignalData.Signals {
		fmt.Println("Testing", signal.Id)
		evocation := sc.Evoke([]base.DistributedSignal{base.NewDistributedSignal("void")}, signal.ToDistributedSignal())
		for _, sig := range evocation {
			fmt.Println("SuperCluster:\t", sig.Represent())
		}
	}
}
