package cmds

import "fmt"
import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"
import "sharedkernel"

type TuneToStationCmd struct {
	stationID sharedkernel.StationID
}

func NewTuneToStationCmd(stationID sharedkernel.StationID) TuneToStationCmd {
	return TuneToStationCmd{stationID}
}

func (cmd TuneToStationCmd) Execute(state *domain.TunerState, ports *infrastructure.OuterWorldPorts) {
	// check business rule
	if state.Subscription == false {
		fmt.Printf("TuneToStationCmd.Execute: cant tune to station %v, subscription inactive\n", cmd.stationID)
		return
	}

	// check business rule
	if int(cmd.stationID) >= len(state.Stations) {
		fmt.Printf("TuneToStationCmd.Execute: cant tune to station %v, no such station\n", cmd.stationID)
		return
	}

	// do!
	ports.HwPort.TuneToStation(cmd.stationID)
}
