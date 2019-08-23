package infrastructure

import "hexagons/tuner/domain"

// TunerServicePort allows the outer world talk to Tuner
type TunerServicePort interface {
	UpdateSubscription(subscription domain.Subscription)
	UpdateStationList(stationList domain.StationList)
	TuneToStation(stationID domain.StationID)
}
