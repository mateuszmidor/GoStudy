package ui

import "time"
import "hexagons/ui/domain"
import "hexagons/ui/application"
import "hexagons/ui/infrastructure"

type UiRoot struct {
	ui domain.Ui
	service application.UiService
	ports infrastructure.Ports
}

func NewUiRoot() UiRoot {
	root := UiRoot{}
	root.ui = domain.NewUi()
	root.service = application.NewUiService(&root.ui)
	root.ports = infrastructure.Ports{root.service, nil}
	return root
}

func (ui *UiRoot) SetupTunerPortOut(tuner infrastructure.TunerPortOut) {
	ui.ports.TunerPortOut = tuner
}

func (ui *UiRoot) GetTunerPortIn() infrastructure.TunerPortIn {
	return ui.ports.TunerPortIn 
}

func (root *UiRoot) Run() {
	for {
		select {
		case <-time.After(5 * time.Second):
			root.ports.TunerPortOut.TuneToStation(application.RandomStation(root.ui.StationList))
		}
	}
}