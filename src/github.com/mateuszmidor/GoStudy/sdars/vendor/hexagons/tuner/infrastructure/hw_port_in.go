package infrastructure

import "hexagons/tuner/domain"

// what Hardware can tell to tuner
type HwPortIn interface {
	SubscriptionUpdated(subscription domain.Subscription)
	StationListUpdated(stationList domain.StationList)
}
