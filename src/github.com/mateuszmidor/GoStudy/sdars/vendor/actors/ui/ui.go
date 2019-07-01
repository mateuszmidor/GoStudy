package ui

import "time"
import "fmt"
import "actors/ui/application"

type UiActor struct {
	stationList []string
	subscriptionActive bool

	// outcoing commands to be assigned with handlers
	OnTuneToStation func(stationId uint32)
}

func NewUiActor() UiActor {
	return UiActor{}
}

// incoming commands
func (a *UiActor) UpdateStationList(stationList []string) {
	fmt.Printf("UiActor.UpdateStationList: %v\n", stationList)
	a.stationList = stationList
}	

// incoming commands
func (a *UiActor) UpdateSubscription(active bool) {
	fmt.Printf("UiActor.UpdateSubscription: %v\n", active)
	a.subscriptionActive = active
}

func (a *UiActor) Run() {
	for {
		select {
		case <-time.After(5 * time.Second):
			a.OnTuneToStation(application.RandomStation(a.stationList))
		}
	}
}