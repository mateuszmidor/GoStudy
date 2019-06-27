package adapters

import (
	"sdars"
	"actors/cluster"
	"cmds"
)

type ClusterAdapter struct {
	commandQueue *sdars.CommandQueue
	a *cluster.ClusterActor
}

func NewClusterAdapter(c *sdars.CommandQueue, a *cluster.ClusterActor) ClusterAdapter {
	ca := ClusterAdapter{c, a}
	a.OnTuneToStation = ca.TuneToStation
	return ca
}

// Transform Cluster command into Tuner command
func (ca *ClusterAdapter) TuneToStation(stationId uint32) {
	*ca.commandQueue <- *cmds.NewTuneToStationCmd(stationId)
}

// Transform Tuner command into Cluster command
func (ca ClusterAdapter) UpdateStationList(stationList []string) {
	// forward to cluster actor
	ca.a.UpdateStationList(stationList)
}

// Transform Tuner command into Cluster command
func (ca ClusterAdapter) UpdateSubscription(active bool) {
	// forward to cluster actor
	ca.a.UpdateSubscription(active)	
}