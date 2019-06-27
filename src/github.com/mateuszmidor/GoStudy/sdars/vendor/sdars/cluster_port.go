package sdars

// what operations the Cluster allows
type ClusterPort interface {
	UpdateStationList(stationList []string)
	UpdateSubscription(active bool)
}