package tests

import (
	"fmt"
	"github.com/Art-of-the-Living/gracious/base"
	"testing"
)

func TestClusterBasic(t *testing.T) {
	testIterations := 5
	cluster := base.NewCluster("test")
	cluster.PassThrough = false
	cluster.CorrelationThreshold = 5
	blank := base.NewDistributedSignal("blank")
	main := base.NewDistributedSignal("testMain")
	main.Features[base.Address{X: 1, Y: 1}] = 1
	main.Features[base.Address{X: 0, Y: 3}] = 1
	associate := base.NewDistributedSignal("testAssc")
	associates := base.NewNetwork("Assc")
	associate.Features[base.Address{X: 0, Y: 2}] = 1
	associate.Features[base.Address{X: 1, Y: 4}] = 1
	// Training and Building Cycles
	for i := 0; i < testIterations; i++ {
		fmt.Println("Iteration Number:", i)
		ds := cluster.Evoke(main, associates.GetConnections("void"))
		fmt.Println(main.Represent())
		fmt.Println(associate.Represent())
		fmt.Println(ds.Represent())
	}
	// Blank Main Signal
	ds := cluster.Evoke(blank, associates.GetConnections("void"))
	fmt.Println(blank.Represent())
	fmt.Println(associate.Represent())
	fmt.Println(ds.Represent())
	// Blank Association Signal
	ds = cluster.Evoke(main, base.NewNetwork("void").GetConnections("void"))
	fmt.Println(main.Represent())
	fmt.Println(blank.Represent())
	fmt.Println(ds.Represent())
}
