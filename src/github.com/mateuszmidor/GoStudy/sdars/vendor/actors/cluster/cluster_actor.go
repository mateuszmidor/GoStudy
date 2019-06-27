package cluster

import "time"
import "fmt"

type ClusterActor struct {
	stationList []string
	subscriptionActive bool

	// outcoing commands to be assigned with handlers
	OnTuneToStation func(stationId uint32)
}

func NewClusterActor() ClusterActor {
	return ClusterActor{}
}

// incoming commands
func (a *ClusterActor) UpdateStationList(stationList []string) {
	fmt.Printf("ClusterActor.UpdateStationList: %v\n", stationList)
	a.stationList = stationList
}	

// incoming commands
func (a *ClusterActor) UpdateSubscription(active bool) {
	fmt.Printf("ClusterActor.UpdateSubscription: %v\n", active)
	a.subscriptionActive = active
}

func (a *ClusterActor) Run() {
	for {
		select {
		case <-time.After(5 * time.Second):
			a.OnTuneToStation(RandomStation(a.stationList))
		}
	}
}