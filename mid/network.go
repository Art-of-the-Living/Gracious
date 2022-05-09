package mid

import "github.com/Art-of-the-Living/gracious/base"

// A Network is a collection of SubNet instances. A Network permits the reading of SubNets and
// allows Signals to be added to a subnet as well as to retrieve a subnet.
type Network struct {
	name       string
	subnetwork map[string]*SubNet
}

func NewNetwork(name string) *Network {
	net := Network{name: "Network:" + name, subnetwork: make(map[string]*SubNet)}
	return &net
}

func (n *Network) AddSignals(destinationId string, signal ...base.DistributedSignal) {
	for _, sig := range signal {
		if _, ok := n.subnetwork[destinationId]; !ok {
			n.subnetwork[destinationId] = NewSubNet()
		}
		n.subnetwork[destinationId].Add(sig)
	}
}

func (n *Network) GetSubNet(destinationId string) *SubNet {
	if _, ok := n.subnetwork[destinationId]; !ok {
		n.subnetwork[destinationId] = NewSubNet()
	}
	return n.subnetwork[destinationId]
}

// WinnersTakeAll is similar to the DistributedSignal WinnersTakeAll method, but when
// called on a Network it forces each and every DistributedSignal to undergo a
// WinnersTakeAll function.
func (n *Network) WinnersTakeAll(gap int) {
	for _, connector := range n.subnetwork {
		for _, signal := range connector.Signals {
			signal.WinnersTakeAll(gap)
		}
	}
}

type SubNet struct {
	Signals map[string]base.DistributedSignal
}

func NewSubNet(signal ...base.DistributedSignal) *SubNet {
	c := SubNet{Signals: make(map[string]base.DistributedSignal)}
	for _, sig := range signal {
		c.Add(sig)
	}
	return &c
}

func (sn *SubNet) Add(signal base.DistributedSignal) {
	sn.Signals[signal.Id] = signal
}
