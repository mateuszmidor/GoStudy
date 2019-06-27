package hardware

import "time"
import "fmt"
import "math/rand"

type HwActor struct {
	currentStationId uint32

	// outcoing commands to be assigned with handlers
	OnUpdatStationList func([]string)
	OnUpdateSubscription func(bool)
}

func NewHwActor() HwActor {
	return HwActor{}
}

// incoming commands
func (a *HwActor) TuneToStation(stationId uint32) {
	fmt.Printf("HwActor.TuneToStation: %v\n", stationId)
	a.currentStationId = stationId
}	

func (a *HwActor) Run() {
	// activate subscription
	a.OnUpdateSubscription(true)

	// send hw events...
	for {
		time.Sleep(3000 * time.Millisecond)
		switch lucky := rand.Intn(10); lucky {
		case 1,2,3:
			a.OnUpdatStationList(RandomStationList())
		case 4:
			a.OnUpdateSubscription(RandomSubscription())
		default:
			;
		}
	}
}