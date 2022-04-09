package console

import "gracious/model"

func ASCIIToQuale(character rune) model.Quale {
	quale := model.NewQuale(256)
	quale.SetFeature(int(character), 1)
	return quale
}

// The ReadConsole grants access to a neuron group for communicating sensory information. A read console receives
// external console characters as representations of traditional ASCII characters directly into the system without
// sensory pre-processing of simulated or imported external data.
type ReadConsole struct {
	ImportingChannel *model.Quale
}
