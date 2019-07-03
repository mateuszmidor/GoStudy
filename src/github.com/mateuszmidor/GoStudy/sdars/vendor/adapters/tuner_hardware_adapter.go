package adapters

import (
	"actors/hardware"
	"hexagons/tuner"
	"hexagons/tuner/domain"
	tunerports "hexagons/tuner/infrastructure"
)

// Implements tuner: HwPortOut
type HardwareAdapter struct {
	tunerPortIn tunerports.HwPortIn
	a *hardware.HwActor
}

func NewHardwareAdapter(r *tuner.TunerRoot, a *hardware.HwActor) HardwareAdapter {
	ha := HardwareAdapter{r.GetHwPortIn(), a}
	a.OnUpdatStationList = ha.UpdateStationList
	a.OnUpdateSubscription = ha.UpdateSubscription
	return ha
}

// Transform HW command into Tuner command
func (ha *HardwareAdapter) UpdateStationList(newStationList domain.StationList) {
	ha.tunerPortIn.StationListUpdated(newStationList)
}

// Transform HW command into Tuner command
func (ha *HardwareAdapter) UpdateSubscription(subscription domain.Subscription) {
	ha.tunerPortIn.SubscriptionUpdated(subscription)
}

// Transform Tuner command into HW command
func (ha *HardwareAdapter) TuneToStation(stationId uint32) {
	// forward to HW actor
	ha.a.TuneToStation(stationId)
}
