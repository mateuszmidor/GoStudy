package application

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"
import "hexagons/tuner/application/cmds"

// Implements Tuner input ports: UiPortIn, HwPortIn
type TunerService struct {
	commandQueue CommandQueue
}

func NewTunerService() TunerService {
	return TunerService{NewCommandQueue()}
}

func (service *TunerService) putCommand(cmd Cmd) {
	service.commandQueue <- cmd
}

// UiPortIn
func (service *TunerService) TuneToStation(stationId domain.StationId) {
	service.putCommand(cmds.NewTuneToStationCmd(stationId))
}

// HwPortIn
func (service *TunerService) SubscriptionUpdated(subscription domain.Subscription) {
	service.putCommand(cmds.NewUpdateSubscriptionCmd(subscription))
}

// HwPortIn
func (service *TunerService) StationListUpdated(stationList domain.StationList) {
	service.putCommand(cmds.NewUpdateStationListCmd(stationList))
}

// To be run from non-main gorutine
func (service *TunerService) Run(tuner *domain.Tuner, ports *infrastructure.Ports) {
	// loop forever
	for {
		select {
		case cmd:= <- service.commandQueue:
			cmd.Execute(tuner, ports)
		}
	}
}