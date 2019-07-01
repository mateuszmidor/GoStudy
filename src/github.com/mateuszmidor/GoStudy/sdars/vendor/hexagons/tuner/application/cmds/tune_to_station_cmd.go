package cmds

import "fmt"
import "hexagons/tuner"
import "hexagons/tuner/domain"

type TuneToStationCmd struct {
	root *tuner.TunerRoot
	stationId domain.StationId
}

func NewTuneToStationCmd(root *tuner.TunerRoot, stationId domain.StationId) *TuneToStationCmd {
	return &TuneToStationCmd{root, stationId}
}

func (cmd TuneToStationCmd) Execute() {
	// check business rule
	if cmd.root.Tuner.Subscription == false {
		fmt.Printf("TuneToStationCmd.Execute: cant tune, subscription inactive\n")
		return
	}

	if int(cmd.stationId) >= len(cmd.root.Tuner.Stations) {
		fmt.Printf("TuneToStationCmd.Execute: cant tune to station %v, no such station\n", cmd.stationId)
		return	
	}
	
	cmd.root.HardwarePortOut.TuneToStation(cmd.stationId)
}