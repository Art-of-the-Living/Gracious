package tests

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/mid"
	"github.com/Art-of-the-Living/gracious/tests/util"
	"testing"
)

func TestBasic(t *testing.T) {
	var sc = mid.NewSuperCluster("test")
	sc.PassThrough = true
	var seq = mid.NewSequencer("test", 8)
	var tr = util.NewTextReader("Testing")
	i := 0
	for tr.Next() {
		fmt.Println("\t Iteration:", i)
		value := seq.Sequence(tr.Evoke())
		for _, sig := range value {
			fmt.Println(sig.Represent())
		}
		evocation := sc.Evoke(value)
		for _, sig := range evocation {
			fmt.Println(sig.Represent())
		}
		i++
	}
}
