package domain

type HwState struct {
	CurrentStationId uint32
}

func NewHwState() HwState {
	return HwState{}
}
