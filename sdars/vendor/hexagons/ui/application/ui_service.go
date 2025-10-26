package application

import "hexagons/ui/domain"
import "hexagons/ui/infrastructure"
import "fmt"
import "time"

// Implements UiServicePort
type UiService struct {
	state *domain.UiState
}

func NewUiService(state *domain.UiState) UiService {
	return UiService{state}
}

// UiServicePort
func (service *UiService) UpdateSubscription(active bool) {
	fmt.Printf("UiService.UpdateSubscription: %v\n", active)
	service.state.SubscriptionActive = active
}

// UiServicePort
func (service *UiService) UpdateStationList(stations []string) {
	fmt.Printf("UiService.UpdateStationList: %v\n", stations)
	service.state.StationList = stations
}

// Generate random Ui commands and send them to Tuner
func (service *UiService) Run(ports *infrastructure.OuterWorldPorts) {
	for {
		select {
		case <-time.After(5 * time.Second):
			ports.TunerPort.TuneToStation(RandomStation(service.state.StationList))
		}
	}
}
