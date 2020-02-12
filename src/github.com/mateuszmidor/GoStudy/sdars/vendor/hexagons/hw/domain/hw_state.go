package domain

import "sharedkernel"

type HwState struct {
	CurrentStationId sharedkernel.StationID
}

func NewHwState() HwState {
	return HwState{}
}
