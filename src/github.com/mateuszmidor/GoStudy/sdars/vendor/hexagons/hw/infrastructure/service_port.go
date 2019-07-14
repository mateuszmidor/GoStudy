package infrastructure

type ServicePort interface {
	TuneToStation(stationId uint32)
}
