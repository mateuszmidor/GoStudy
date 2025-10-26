package utils

type Result struct {
	msg string
}

func NewResult(msg string) Result {
	return Result{msg}
}
