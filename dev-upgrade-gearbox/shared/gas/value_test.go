package gas_test

import (
	"shared/gas"
	"testing"
)

func TestNewPanicsWhenBelowValidRange(t *testing.T) {
	defer func() {
		p := recover()
		if p == nil {
			t.Errorf("gas.New(-1) should panic")
		}
	}()
	_ = gas.New(-1)
}

func TestNewPanicsWhenAboveValidRange(t *testing.T) {
	defer func() {
		p := recover()
		if p == nil {
			t.Errorf("gas.New(1.1) should panic")
		}
	}()
	_ = gas.New(1.1)
}
