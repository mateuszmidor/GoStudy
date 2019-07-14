package infrastructure

import "hexagons/tuner/domain"

// what services the Tuner offers
type ServicePort interface {
	SubscriptionUpdated(subscription domain.Subscription)
	StationListUpdated(stationList domain.StationList)
	TuneToStation(stationId domain.StationId)
}
