package gearadjustment_test

import (
	"driver/gearadjustment"
	"shared/gear"
	"testing"
)

func TestGearDownWhenTrailorAndDownSlope(t *testing.T) {
	// given
	isTrailorAttached := true
	isDownSlope := true

	// when
	gc := gearadjustment.AdjustGearForEngineBreaking(isDownSlope, isTrailorAttached)

	// then
	if gc != gear.GearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}

func TestKeepGearWhenNoTrailorDownSlope(t *testing.T) {
	// given
	isTrailorAttached := false
	isDownSlope := true

	// when
	gc := gearadjustment.AdjustGearForEngineBreaking(isDownSlope, isTrailorAttached)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}

func TestKeepGearWhenTrailorNoDownSlope(t *testing.T) {
	// given
	isTrailorAttached := true
	isDownSlope := false

	// when
	gc := gearadjustment.AdjustGearForEngineBreaking(isDownSlope, isTrailorAttached)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}
