package cmds

import "fmt"
import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type TuneToStationCmd struct {
	stationId domain.StationId
}

func NewTuneToStationCmd(stationId domain.StationId) *TuneToStationCmd {
	return &TuneToStationCmd{stationId}
}

func (cmd TuneToStationCmd) Execute(tuner *domain.Tuner, ports *infrastructure.Ports) {
	// check business rule
	if tuner.Subscription == false {
		fmt.Printf("TuneToStationCmd.Execute: cant tune, subscription inactive\n")
		return
	}

	if int(cmd.stationId) >= len(tuner.Stations) {
		fmt.Printf("TuneToStationCmd.Execute: cant tune to station %v, no such station\n", cmd.stationId)
		return	
	}
	
	ports.HardwarePortOut.TuneToStation(cmd.stationId)
}