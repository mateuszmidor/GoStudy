package infrastructure

// what UI can tell to tuner
type TunerCommandsPort interface {
	TuneToStation(stationId uint32)
}
