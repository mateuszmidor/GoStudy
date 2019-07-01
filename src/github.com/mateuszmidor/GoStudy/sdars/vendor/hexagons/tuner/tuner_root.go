package tuner

import "hexagons/tuner/domain"
import "hexagons/tuner/application"
import "hexagons/tuner/infrastructure"

// Tuner aggregate root; visible to the outer world
type TunerRoot struct {
	HardwarePortOut infrastructure.HardwarePortOut
	GuiPortOut infrastructure.GuiPortOut
	CommandQueue application.CommandQueue
	Tuner domain.Tuner
}

func NewTunerRoot() TunerRoot {
	return TunerRoot{nil, nil, application.NewCommandQueue(), domain.NewTuner()}
}

func (t *TunerRoot) SetupPorts(hardwarePortOut infrastructure.HardwarePortOut, guiPortOut infrastructure.GuiPortOut) {
	t.HardwarePortOut = hardwarePortOut
	t.GuiPortOut = guiPortOut
}

// To be run from non-main gorutine
func (t *TunerRoot) Run() {
	// loop forever
	for {
		select {
		case cmd:= <- t.CommandQueue:
			cmd.Execute()
		}
	}
}