package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New can't return nil!")
	} else {
		tracer.Trace("Hello in trace package")
		if buf.String() != "Hello in trace package\n" {
			t.Errorf("Trace method generated wrong string: '%s'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("trace me")
}
