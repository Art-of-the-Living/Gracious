package base

// A Network is a collection of like DistributedSignals. A Network permits the reading
// of distributed signals by different parts of an architecture without the need to
// manage DistributedSignal maps directly.
type Network struct {
	name       string
	connectors map[string]*Connections
}

type Connections struct {
	Signals map[string]DistributedSignal
}

func NewConnector(signal ...DistributedSignal) *Connections {
	c := Connections{Signals: make(map[string]DistributedSignal)}
	for _, sig := range signal {
		c.Add(sig)
	}
	return &c
}

func (c *Connections) Add(signal DistributedSignal) {
	c.Signals[signal.Name] = signal
}

func NewNetwork(name string) *Network {
	net := Network{name: "Network:" + name, connectors: make(map[string]*Connections)}
	return &net
}

func (n *Network) AddSignals(destinationId string, signal ...DistributedSignal) {
	for _, sig := range signal {
		if _, ok := n.connectors[destinationId]; !ok {
			n.connectors[destinationId] = NewConnector()
		}
		n.connectors[destinationId].Add(sig)
	}
}

func (n *Network) GetConnections(destinationId string) *Connections {
	if _, ok := n.connectors[destinationId]; !ok {
		n.connectors[destinationId] = NewConnector()
	}
	return n.connectors[destinationId]
}

// WinnersTakeAll is similar to the DistributedSignal WinnersTakeAll method, but when
// called on a Network it forces each and every DistributedSignal to undergo a
// WinnersTakeAll function.
func (n *Network) WinnersTakeAll(gap int) {
	for _, connector := range n.connectors {
		for _, signal := range connector.Signals {
			signal.WinnersTakeAll(gap)
		}
	}
}
