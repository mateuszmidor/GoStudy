package cmds

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type UpdateStationListCmd struct {
	stations domain.StationList
}

func NewUpdateStationListCmd(stations domain.StationList) UpdateStationListCmd {
	return UpdateStationListCmd{stations}
}

func (cmd UpdateStationListCmd) Execute(tuner *domain.Tuner, ports *infrastructure.Ports) {
	tuner.Stations = cmd.stations
	ports.UiPortOut.UpdateStationList(cmd.stations)
}
