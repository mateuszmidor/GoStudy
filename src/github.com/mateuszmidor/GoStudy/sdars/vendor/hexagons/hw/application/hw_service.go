package application

import "hexagons/hw/domain"
import "hexagons/hw/infrastructure"
import "fmt"
import "time"
import "math/rand"

// Implements input pots: TunerInPort
type HwService struct {
	hw *domain.Hw
}

func NewHwService(hw *domain.Hw) HwService {
	return HwService{hw}
}

// TunerInPort
func (service *HwService) TuneToStation(stationId uint32) {
	fmt.Printf("HwService.TuneToStation: %v\n", stationId)
	service.hw.CurrentStationId = stationId
}

// Generate random Hw events and send them to Tuner port
func (service *HwService) Run(ports infrastructure.Ports) {
	// activate subscription
	ports.TunerOutPort.UpdateSubscription(true)

	// send hw events...
	for {
		time.Sleep(3000 * time.Millisecond)
		switch lucky := rand.Intn(10); lucky {
		case 1,2,3:
			ports.TunerOutPort.UpdateStationList(RandomStationList())
		case 8:
			ports.TunerOutPort.UpdateSubscription(RandomSubscription())
		default:
			;
		}
	}
}