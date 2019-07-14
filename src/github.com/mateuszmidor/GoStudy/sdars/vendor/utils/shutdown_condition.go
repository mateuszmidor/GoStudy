package utils

import "os"
import "os/signal"
import "syscall"

type ShutdownCondition struct {
	shutdownRequested chan os.Signal
}

func NewShutdownCondition() *ShutdownCondition {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return &ShutdownCondition{c}
}

func (c *ShutdownCondition) Wait() {
	<-c.shutdownRequested
}
