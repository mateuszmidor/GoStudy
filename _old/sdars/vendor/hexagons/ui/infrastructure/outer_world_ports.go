package infrastructure

type TunerPort interface {
	TuneToStation(stationId uint32)
}

type OuterWorldPorts struct {
	TunerPort TunerPort
}
