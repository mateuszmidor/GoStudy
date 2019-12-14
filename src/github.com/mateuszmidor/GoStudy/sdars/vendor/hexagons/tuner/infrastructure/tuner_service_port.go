package infrastructure

import "sharedkernel"

// TunerServicePort allows the outer world talk to Tuner
type TunerServicePort interface {
	UpdateSubscription(subscription sharedkernel.Subscription)
	UpdateStationList(stationList sharedkernel.StationList)
	TuneToStation(stationID sharedkernel.StationID)
}
