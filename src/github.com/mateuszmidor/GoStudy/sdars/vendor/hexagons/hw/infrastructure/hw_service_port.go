package infrastructure

type HwServicePort interface {
	TuneToStation(stationId uint32)
}
