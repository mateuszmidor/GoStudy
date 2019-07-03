package application

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"
import "hexagons/tuner/application/cmds"

// Implements UiPortIn, HwPortIn
type TunerService struct {
	commandQueue CommandQueue
}

func NewTunerService() TunerService {
	return TunerService{NewCommandQueue()}
}

func (service *TunerService) PutCommand(cmd Cmd) {
	service.commandQueue <- cmd
}

func (service TunerService) TuneToStation(stationId domain.StationId) {
	service.PutCommand(cmds.NewTuneToStationCmd(stationId))
}

func (service TunerService) SubscriptionUpdated(subscription domain.Subscription) {
	service.PutCommand(cmds.NewUpdateSubscriptionCmd(subscription))
}

func (service TunerService) StationListUpdated(stationList domain.StationList) {
	service.PutCommand(cmds.NewUpdateStationListCmd(stationList))
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