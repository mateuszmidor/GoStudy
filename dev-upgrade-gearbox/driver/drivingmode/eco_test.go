package drivingmode_test

import (
	"driver/aggressiveness"
	"driver/drivingmode"
	"shared/gear"
	"driver/types"
	"driver/util"
	"shared/gas"
	"testing"
)

func TestGetOptimalRange(t *testing.T) {
	// given
	var aggressiveness aggressiveness.Level = nil // aggressiveness not used in Eco mode
	var eco drivingmode.Mode = drivingmode.NewEco(1000.0, 2000.0)

	// when
	var minRPM, maxRPM types.RPM = eco.GetOptimalRPM(aggressiveness)

	// then
	if !util.IsEqual(1000.0, minRPM) {
		t.Errorf("Expected minRPM %f, got %f", 1000.0, minRPM)
	}

	if !util.IsEqual(2000.0, maxRPM) {
		t.Errorf("Expected maxRPM %f, got %f", 2000.0, maxRPM)
	}
}

func TestKickDown1ReturnsKeepGear(t *testing.T) {
	// given
	var eco drivingmode.Mode = drivingmode.NewEco(1000.0, 2000.0)

	// when
	gc := eco.GetKickDownGearChange(gas.Full)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}

func TestKickDown2ReturnsKeepGear(t *testing.T) {
	// given
	var eco drivingmode.Mode = drivingmode.NewEco(1000.0, 2000.0)

	// when
	gc := eco.GetKickDownGearChange(gas.Full)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}
