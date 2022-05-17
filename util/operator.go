package util

// An Operator is a non-neuronal circuit for processing QualitativeSignal values. The Execute function is evoked
// to perform the given operation of the Operator.
type Operator interface {
	Execute(signal QualitativeSignal) QualitativeSignal
}

// WinnerTakesAll forces the Features in the QualitativeSignal to fight for dominance and only the strongest features
// will remain present in the signal. The gap parameter permits a level of tolerance for features which almost meet
// with max threshold. No signal beneath 1 will ever be passed through. Signals with values above 4, will be reduced
// by half.
type WinnerTakesAll struct {
	Gap int
}

// NewWTA returns a new WinnerTakesAll instance
func NewWTA(gap int) *WinnerTakesAll {
	wta := WinnerTakesAll{Gap: gap}
	return &wta
}

// Execute implements the Execute function of the Operator interface.
func (wta *WinnerTakesAll) Execute(signal QualitativeSignal) QualitativeSignal {
	signal.WinnerTakesAll(wta.Gap)
	return signal
}
