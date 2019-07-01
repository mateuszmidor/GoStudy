package application

type CommandQueue chan Cmd

func NewCommandQueue() CommandQueue {
	return make(CommandQueue, 100)
}