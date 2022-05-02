package tests

import (
	"github.com/Art-of-the-Living/gracious/base"
	"testing"
)

func TestDistributedSignal(t *testing.T) {
	ds := base.NewDistributedSignal("test")
	println(ds.Represent())
	testValue := 3
	testAddr := base.Address{X: 1, Y: 1}
	ds.Features[testAddr] = testValue
	println(ds.Represent())
	if ds.Features[testAddr] != testValue {
		panic("Basic address-value retention is failing!")
	}

}
