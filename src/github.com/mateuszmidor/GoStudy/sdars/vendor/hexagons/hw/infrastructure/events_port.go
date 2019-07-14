package infrastructure

type EventsPort interface {
	UpdateStationList(stations []string)
	UpdateSubscription(active bool)
}
