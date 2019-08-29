package domain

// TunerState - internal state of the tuner
type TunerState struct {
	Stations     StationList
	Subscription Subscription
}

func NewTunerState() TunerState {
	return TunerState{}
}
