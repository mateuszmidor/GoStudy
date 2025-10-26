package application

import "hexagons/hw/domain"
import "hexagons/hw/infrastructure"
import "fmt"
import "time"
import "math/rand"

// Implements HwServicePort
type HwService struct {
	state *domain.HwState
}

func NewHwService(state *domain.HwState) HwService {
	return HwService{state}
}

// HwServicePort
func (service *HwService) TuneToStation(stationId uint32) {
	fmt.Printf("HwService.TuneToStation: %v\n", stationId)
	service.state.CurrentStationId = stationId
}

// Generate random Hw events and send them to Tuner port
func (service *HwService) Run(ports *infrastructure.OuterWorldPorts) {
	// activate subscription
	ports.TunerPort.UpdateSubscription(true)

	// send hw events...
	for {
		time.Sleep(3000 * time.Millisecond)
		switch lucky := rand.Intn(10); lucky {
		case 1, 2, 3:
			ports.TunerPort.UpdateStationList(RandomStationList())
		case 8:
			ports.TunerPort.UpdateSubscription(RandomSubscription())
		default:

		}
	}
}
