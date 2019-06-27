package sdars

import "fmt"

type TunerCommandProcessor struct {
	Tuner
	CommandQueue CommandQueue
}

func NewTunerCommandProcessor() TunerCommandProcessor {
	return TunerCommandProcessor{NewTuner(), NewCommandQueue()}
}

// func (c *TunerCommandProcessor) PutCommand(cmd Cmd) {
// 	c.commandQueue <- &cmd
// }

func (c *TunerCommandProcessor) Run(shutdownCondition ShutdownCondition) {
	for {
		select {
		case cmd:= <- c.CommandQueue:
			cmd.Execute(&c.Tuner)
		case <-shutdownCondition.ShutdownRequested:
			fmt.Printf("CommandProcessor.Run: requested shutdown\n")
			return
		}
	}
}