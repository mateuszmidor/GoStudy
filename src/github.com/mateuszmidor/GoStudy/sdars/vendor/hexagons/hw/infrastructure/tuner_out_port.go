package infrastructure

type TunerOutPort interface {
	UpdateStationList(stations []string)
	UpdateSubscription(active bool)
}