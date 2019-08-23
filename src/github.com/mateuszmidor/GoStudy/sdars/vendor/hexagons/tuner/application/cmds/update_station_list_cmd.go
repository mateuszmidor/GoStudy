package cmds

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type UpdateStationListCmd struct {
	stations domain.StationList
}

func NewUpdateStationListCmd(stations domain.StationList) UpdateStationListCmd {
	return UpdateStationListCmd{stations}
}

func (cmd UpdateStationListCmd) Execute(state *domain.TunerState, ports *infrastructure.OuterWorldPorts) {
	state.Stations = cmd.stations
	ports.UIPort.UpdateStationList(cmd.stations)
}
