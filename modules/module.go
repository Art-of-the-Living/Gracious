package modules

import "github.com/KennethGrace/gracious/model"

// A Module is a cluster of neuron groups which processes input or produces output. Input is processed via registered
// phenomena. Output is produced via registered actions.
type Module interface {
	RegisterPhenomena(phenomena model.Phenomena) error
}
