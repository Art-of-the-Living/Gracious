package objects

import (
	"github.com/Art-of-the-Living/gracious/base"
)

type Evokable interface {
	Evoke(aSignals []base.DistributedSignal, mSignals ...base.DistributedSignal) []base.DistributedSignal
}
