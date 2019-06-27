package cmds

import "sdars"

type UpdateStationListCmd struct {
	stationList []string
}

func NewUpdateStationListCmd(stationList []string) *UpdateStationListCmd {
	return &UpdateStationListCmd{stationList}
}

func (cmd UpdateStationListCmd) Execute(tuner *sdars.Tuner) {
	tuner.State.StationList = cmd.stationList
	tuner.ClusterPort.UpdateStationList(cmd.stationList)
}