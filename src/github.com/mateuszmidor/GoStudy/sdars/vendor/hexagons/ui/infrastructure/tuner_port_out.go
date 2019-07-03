package infrastructure

// what UI can tell to tuner
type TunerPortOut interface {
	TuneToStation(stationId uint32)
}