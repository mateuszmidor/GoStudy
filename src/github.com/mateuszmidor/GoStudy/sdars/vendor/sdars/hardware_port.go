package sdars

// what operations the Hardware allows
type HardwarePort interface {
	TuneToStation(stationId uint32)
}