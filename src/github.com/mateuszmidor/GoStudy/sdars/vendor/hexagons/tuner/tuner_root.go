package tuner

import "hexagons/tuner/domain"
import "hexagons/tuner/application"
import "hexagons/tuner/infrastructure"

// Tuner aggregate root; visible to the outer world
type TunerRoot struct {
	Tuner domain.Tuner
	Ports infrastructure.Ports
	CommandQueue application.CommandQueue
}

func NewTunerRoot() TunerRoot {
	return TunerRoot{domain.NewTuner(), infrastructure.Ports{}, application.NewCommandQueue()}
}

func (t *TunerRoot) SetupPorts(hardwarePortOut infrastructure.HardwarePortOut, guiPortOut infrastructure.GuiPortOut) {
	t.Ports.HardwarePortOut = hardwarePortOut
	t.Ports.GuiPortOut = guiPortOut
}

// To be run from non-main gorutine
func (t *TunerRoot) Run() {
	// loop forever
	for {
		select {
		case cmd:= <- t.CommandQueue:
			cmd.Execute(&t.Tuner, &t.Ports)
		}
	}
}