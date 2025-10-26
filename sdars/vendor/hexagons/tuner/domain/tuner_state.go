package domain

import "sharedkernel"

// TunerState - internal state of the tuner
type TunerState struct {
	Stations     sharedkernel.StationList
	Subscription sharedkernel.Subscription
}

// NewTunerState - constructor
func NewTunerState() TunerState {
	return TunerState{}
}
