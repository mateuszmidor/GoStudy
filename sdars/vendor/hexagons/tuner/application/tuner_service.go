package application

import (
	"hexagons/tuner/application/cmds"
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"sharedkernel"
)

// Implements TunerServicePort
type TunerService struct {
	commandQueue CommandQueue
}

func NewTunerService() TunerService {
	return TunerService{NewCommandQueue()}
}

func (service *TunerService) putCommand(cmd Cmd) {
	service.commandQueue <- cmd
}

// TunerServicePort
func (service *TunerService) TuneToStation(stationId sharedkernel.StationID) {
	service.putCommand(cmds.NewTuneToStationCmd(stationId))
}

// TunerServicePort
func (service *TunerService) UpdateSubscription(subscription sharedkernel.Subscription) {
	service.putCommand(cmds.NewUpdateSubscriptionCmd(subscription))
}

// TunerServicePort
func (service *TunerService) UpdateStationList(stationList sharedkernel.StationList) {
	service.putCommand(cmds.NewUpdateStationListCmd(stationList))
}

// To be run from non-main gorutine
func (service *TunerService) Run(state *domain.TunerState, ports *infrastructure.OuterWorldPorts) {
	// loop forever
	for {
		select {
		case cmd := <-service.commandQueue:
			cmd.Execute(state, ports)
		}
	}
}
