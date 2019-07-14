package application

import "hexagons/ui/domain"
import "hexagons/ui/infrastructure"
import "fmt"
import "time"

// Implements ServicePort
type UiService struct {
	ui *domain.Ui
}

func NewUiService(ui *domain.Ui) UiService {
	return UiService{ui}
}

// ServicePort
func (service *UiService) UpdateSubscription(active bool) {
	fmt.Printf("UiService.UpdateSubscription: %v\n", active)
	service.ui.SubscriptionActive = active
}

// ServicePort
func (service *UiService) UpdateStationList(stations []string) {
	fmt.Printf("UiService.UpdateStationList: %v\n", stations)
	service.ui.StationList = stations
}

// Generate random Ui commands and send them to Tuner
func (service *UiService) Run(ports *infrastructure.Ports) {
	for {
		select {
		case <-time.After(5 * time.Second):
			ports.CommandsPort.TuneToStation(RandomStation(service.ui.StationList))
		}
	}
}
