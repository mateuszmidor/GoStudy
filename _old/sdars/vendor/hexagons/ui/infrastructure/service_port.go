package infrastructure

// what services UI provides?
type UiServicePort interface {
	UpdateSubscription(active bool)
	UpdateStationList(stations []string)
}
