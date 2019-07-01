package cmds

import "hexagons/tuner"
import "hexagons/tuner/domain"

type UpdateStationListCmd struct {
	root *tuner.TunerRoot
	stations domain.StationList
}

func NewUpdateStationListCmd(root *tuner.TunerRoot, stations domain.StationList) *UpdateStationListCmd {
	return &UpdateStationListCmd{root, stations}
}

func (cmd UpdateStationListCmd) Execute() {
	cmd.root.Tuner.Stations = cmd.stations
	cmd.root.GuiPortOut.UpdateStationList(cmd.stations)
}