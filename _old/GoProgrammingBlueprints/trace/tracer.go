package trace

import (
	"fmt"
	"io"
)

// Tracer interface allows for tracing code execution
type Tracer interface {
	Trace(...interface{})
}

// tracer is actual Tracer interface implementation
type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// Tracer interface implementation that is silent ie. does not trace
type niltracer struct{}

func (t *niltracer) Trace(a ...interface{}) {}

// New creates new real tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off creates new silent tracer
func Off() Tracer {
	return &niltracer{}
}
