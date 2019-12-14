package cmds

import (
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"sharedkernel"
)

type UpdateStationListCmd struct {
	stations sharedkernel.StationList
}

func NewUpdateStationListCmd(stations sharedkernel.StationList) UpdateStationListCmd {
	return UpdateStationListCmd{stations}
}

func (cmd UpdateStationListCmd) Execute(state *domain.TunerState, ports *infrastructure.OuterWorldPorts) {
	state.Stations = cmd.stations
	ports.UIPort.UpdateStationList(cmd.stations)
}
