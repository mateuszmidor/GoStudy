package cmds

import "sdars"
import "fmt"

type TuneToStationCmd struct {
	stationId uint32
}

func NewTuneToStationCmd(stationId uint32) *TuneToStationCmd {
	return &TuneToStationCmd{stationId}
}

func (cmd TuneToStationCmd) Execute(tuner *sdars.Tuner) {
	// check business rule
	if tuner.State.ActiveSubscription == false {
		fmt.Printf("TuneToStationCmd.Execute: cant tune, subscription inactive\n")
		return
	}

	if int(cmd.stationId) >= len(tuner.State.StationList) {
		fmt.Printf("TuneToStationCmd.Execute: cant tune to station %v, no such station\n", cmd.stationId)
		return	
	}
	
	tuner.HardwarePort.TuneToStation(cmd.stationId)
}