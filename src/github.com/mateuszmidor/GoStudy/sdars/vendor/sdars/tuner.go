package sdars

// Tuner holds no application/domain logic.
// It just holds the state and connections to the outer world
type Tuner struct {
	HardwarePort HardwarePort
	ClusterPort ClusterPort
	State TunerState
}

func NewTuner() Tuner {
	return Tuner{nil, nil, TunerState{}}
}

func (t *Tuner) SetupPorts(hardwarePort HardwarePort, clusterPort ClusterPort) {
	t.HardwarePort = hardwarePort
	t.ClusterPort = clusterPort
}

func (t *Tuner) ExecuteCommand(cmd Cmd) {
	cmd.Execute(t)
}