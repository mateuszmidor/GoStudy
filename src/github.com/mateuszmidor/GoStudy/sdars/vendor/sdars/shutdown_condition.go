package sdars

import "os"
import "os/signal"
import "syscall"

type ShutdownCondition struct {
	ShutdownRequested chan os.Signal
}

func NewShutdownCondition() ShutdownCondition {
	c := make (chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return ShutdownCondition{c}
}