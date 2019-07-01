package domain

// should all the business logic be here or in commands...?
type Tuner struct {
	Stations StationList
	Subscription Subscription
}

func NewTuner() Tuner {
	return Tuner{}
}