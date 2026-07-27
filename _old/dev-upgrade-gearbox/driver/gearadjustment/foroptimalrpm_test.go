package gearadjustment_test

import (
	"driver/gearadjustment"
	"shared/gear"
	"testing"
	"driver/types"
)

func TestKeepGearWhenRPMisOK(t *testing.T) {
	// given
	var min, max types.RPM = 1000, 2000
	var current types.RPM = 1500

	// when
	gc := gearadjustment.AdjustForOptimalRPM(current, min, max)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected no gear change, got %v", gc)
	}
}

func TestGearUpWhenRPMisTooHigh(t *testing.T) {
	// given
	var min, max types.RPM = 1000, 2000
	var current types.RPM = 2200

	// when
	gc := gearadjustment.AdjustForOptimalRPM(current, min, max)

	// then
	if gc != gear.GearUp {
		t.Errorf("Expected gear up, got %v", gc)
	}
}

func TestGearDownWhenRPMisTooLow(t *testing.T) {
	// given
	var min, max types.RPM = 1000, 2000
	var current types.RPM = 800

	// when
	gc := gearadjustment.AdjustForOptimalRPM(current, min, max)

	// then
	if gc != gear.GearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}
