package infrastructure

// what Tuner can tell to Ui
type TunerPortIn interface {
	UpdateSubscription(active bool)
	UpdateStationList(stations []string)
}