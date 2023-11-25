package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// StackTracer describes error with stack trace from package "github.com/pkg/errors"
type StackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := readFile()
	fmt.Println(err.Error())
	for _, frame := range getStackTrace(err) {
		fmt.Println(frame)
	}
}

func readFile() error {
	return errors.WithMessage(openFile(), "read failed") // errors.Wrap would attach another stack trace, but we only need a message
}

func openFile() error {
	return errors.New("open failed") // errors.New from package "github.com/pkg/errors" attaches stack trace. errors.New from package "errors" does not
}

// utils

func getStackTrace(err error) (result []string) {
	if tracer := extractDeepestStackTrace(err); tracer != nil {
		for _, f := range tracer.StackTrace() {
			bytes, _ := f.MarshalText() // outputs eg. main.openFile /Users/mmidor/SoftwareDevelopment/GoStudy/errors-stacktrace/main.go:45
			frame := string(bytes)
			result = append(result, frame)
		}
	}
	return result
}

func extractDeepestStackTrace(err error) StackTracer {
	var result StackTracer
	for errors.As(err, &result) {
		err = errors.Unwrap(err)
	}
	return result
}
