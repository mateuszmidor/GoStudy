package application

import "hexagons/ui/domain"
import "fmt"

// Implements TunerPortIn
type UiService struct {
	ui *domain.Ui
}

func NewUiService (ui *domain.Ui) UiService {
	return UiService{ui}
}

func (service UiService) UpdateSubscription(active bool) {
	fmt.Printf("UiService.UpdateSubscription: %v\n", active)
	service.ui.SubscriptionActive = active
}

func (service UiService) UpdateStationList(stations []string) {
	fmt.Printf("UiService.UpdateStationList: %v\n", stations)
	service.ui.StationList = stations
}