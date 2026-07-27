package infrastructure

type TunerPort interface {
	UpdateStationList(stations []string)
	UpdateSubscription(active bool)
}

type OuterWorldPorts struct {
	TunerPort TunerPort
}
