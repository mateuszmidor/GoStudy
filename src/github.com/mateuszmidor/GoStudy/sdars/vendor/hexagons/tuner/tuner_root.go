package tuner

import "hexagons/tuner/domain"
import "hexagons/tuner/application"
import "hexagons/tuner/infrastructure"

// Tuner aggregate root; visible to the outer world
type TunerRoot struct {
	tuner domain.Tuner
	ports infrastructure.Ports
	commandQueue application.CommandQueue
}

func NewTunerRoot() TunerRoot {
	return TunerRoot{domain.NewTuner(), infrastructure.Ports{}, application.NewCommandQueue()}
}

func (t *TunerRoot) SetupPorts(hardwarePortOut infrastructure.HardwarePortOut, guiPortOut infrastructure.GuiPortOut) {
	t.ports.HardwarePortOut = hardwarePortOut
	t.ports.GuiPortOut = guiPortOut
}

func (t *TunerRoot) PutCommand(cmd application.Cmd) {
	t.commandQueue <- cmd
}

// To be run from non-main gorutine
func (t *TunerRoot) Run() {
	// loop forever
	for {
		select {
		case cmd:= <- t.commandQueue:
			cmd.Execute(&t.tuner, &t.ports)
		}
	}
}