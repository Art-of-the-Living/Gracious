package mid

import "github.com/Art-of-the-Living/gracious/base"

// A Sequencer is a simple, non-neuronal, operator for recording a series of DistributedSignal instances.
// A Sequencer will produce a slice of DistributedSignal instances, which can later be processed.
//
// The hardest job of a Sequencer is determining the length of a sequence. This value can be set via SetLength,
// but should be determined by some thresholding calculation for differences in incoming DistributedSignals
// as they enter the sequence. This determining function can be set via SetThreshHoldFunction.
type Sequencer struct {
	Id                    string                                                                           // The identifier of this sequencer.
	sequence              []base.DistributedSignal                                                         // The working slice of DistributedSignals
	maxLength             int                                                                              // The hard maximum length of a DistributedSignal
	threshHoldingFunction func(oldSignals []base.DistributedSignal, newSignal base.DistributedSignal) bool // A function for determining if the difference between two DistributedSignals is significant enough to terminate the sequence.
}

func NewSequencer(id string, max int) *Sequencer {
	s := Sequencer{Id: id + "-Seq", sequence: make([]base.DistributedSignal, 0, max), maxLength: max}
	s.threshHoldingFunction = func(oldSignals []base.DistributedSignal, newSignal base.DistributedSignal) bool {
		return newSignal.Count() <= 0
	}
	return &s
}

// SetLength sets the absolute maximum length of a sequence.
func (s *Sequencer) SetLength(max int) {
	s.maxLength = max
}

// SetThreshHoldFunction sets the function that should determine if a significant enough change in signals has occurred
// to terminate the sequence. By default, any empty main signal will terminate the slice.
func (s *Sequencer) SetThreshHoldFunction(function func(oldSignals []base.DistributedSignal, newSignal base.DistributedSignal) bool) {
	s.threshHoldingFunction = function
}

// Sequence on a Sequencer returns a slice of DistributedSignal values. If the sequence hasn't terminated yet, then
// the returned slice is of length 0. Once the sequence has terminated, either by the maximum length or the
// thresholding function, then the Sequence will return the full sequence and reset.
//
// This operation can be compared to the phenomena of being unable to recall extremely long words or sentences.
func (s *Sequencer) Sequence(main base.DistributedSignal) []base.DistributedSignal {
	returnSequence := s.sequence
	if s.threshHoldingFunction(s.sequence, main) {
		s.sequence = make([]base.DistributedSignal, 0, s.maxLength)
	} else {
		if len(s.sequence) < s.maxLength {
			s.sequence = append(s.sequence, main)
		} else {
			returnSequence = s.sequence
			s.sequence = make([]base.DistributedSignal, 0, s.maxLength)
		}
	}
	if len(returnSequence) == 0 {
		returnSequence = []base.DistributedSignal{base.NewDistributedSignal(s.Id)}
	}
	return returnSequence
}

// The DeSequencer performs the opposite role of a sequencer. A DeSequencer receives a slice of DistributedSignal
// values and appends them to an internal queue. DistributedSignals are evoked as they were received, FIFO.
type DeSequencer struct {
}
