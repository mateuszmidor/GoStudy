package infrastructure

// what services do we provide?
type ServicePort interface {
	UpdateSubscription(active bool)
	UpdateStationList(stations []string)
}
