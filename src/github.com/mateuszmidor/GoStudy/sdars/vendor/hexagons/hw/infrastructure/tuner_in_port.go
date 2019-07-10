package infrastructure

type TunerInPort interface {
	TuneToStation(stationId uint32)
}