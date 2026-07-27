package gearboxdriver

import (
	"driver/characteristics"
	"driver/externalsystemsfacade"
	"externalsystems"
	"shared/events"
	"shared/gas"
	"shared/gear"
	"testing"
)

func TestShouldKeepGearWhenRPMisOK(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(1500)
	d := newEcoDriver(ed)

	// when
	gc, _ := d.HandleGas(gas.Half)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep current gear, got %v", gc)
	}
}

func TestShouldGearUpWhenRPMisTooHigh(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(2400)
	d := newEcoDriver(ed)

	// when
	gc, _ := d.HandleGas(gas.Half)

	// then
	if gc != gear.GearUp {
		t.Errorf("Expected gear up, got %v", gc)
	}
}

func TestShouldGearDownWhenRPMisTooLow(t *testing.T) {
	ed := newExternalData().SetRPM(600)
	d := newEcoDriver(ed)

	// when
	gc, _ := d.HandleGas(gas.Half)

	// then
	if gc != gear.GearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}

func TestShouldKeepGearWhenUślizgInMDynamicMode(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(2400).SetUślizg()
	d := newEcoDriver(ed)
	d.SetMDynamic(true)

	// when
	gc, _ := d.HandleGas(gas.Half)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep current gear, got %v", gc)
	}
}

func TestShouldGearDownWhenTrailorAttachedAndDrivingDownTheSlope(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(1500).SetDownTheSlope().SetTrailor()
	d := newEcoDriver(ed)

	// when
	gc, _ := d.HandleGas(gas.Zero)

	// then
	if gc != gear.GearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}

func TestShouldKeepGearWhenEcoOvertaking(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(1500)
	d := newEcoDriver(ed)

	// when
	gc, _ := d.HandleGas(gas.Full)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}

func TestShouldGearDownWhenSlowSportOvertaking(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(1500)
	d := newEcoDriver(ed)
	d.SetDrivingModeSport()
	kickdown1threshold := gas.New(characteristics.GetGasKickDown1Sport())

	// when
	gc, _ := d.HandleGas(kickdown1threshold)

	// then
	if gc != gear.GearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}

func TestShouldDoubleGearDownWhenFastSportOvertaking(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(1500)
	d := newEcoDriver(ed)
	d.SetDrivingModeSport()
	kickdown2threshold := gas.New(characteristics.GetGasKickDown2Sport())

	// when
	gc, _ := d.HandleGas(kickdown2threshold)

	// then
	if gc != gear.DoubleGearDown {
		t.Errorf("Expected gear down, got %v", gc)
	}
}

func TestShouldKeepGearWhenManualGearChangeIsActive(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(2400).SetManualGearChangeActive()
	d := newEcoDriver(ed)

	// when
	gc, _ := d.HandleGas(gas.Half)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep current, got %v", gc)
	}
}

func TestEmitEventGearDownOnGearDown(t *testing.T) {
	// given
	ed := newExternalData().SetRPM(400)
	d := newEcoDriver(ed)

	// when
	_, eventList := d.HandleGas(gas.Half)

	// then
	if !eventList.Contains(events.GearDown) {
		t.Errorf("Expected event gear down, got %v", eventList)
	}
}

func TestShouldKeepGearWithActualExternalSystemsAndStuff(t *testing.T) {
	// given
	lights := externalsystems.NewLights(nil)
	ed := externalsystems.NewExternalSystems(
		1500,
		uslizg.AngularSpeedForNoUślizg,
		lights,
	)
	esFacade := externalsystemsfacade.Facade{
		ExternalSystems: &ed,
	}
	d := newEcoDriver(esFacade)

	// when
	gc, _ := d.HandleGas(gas.Half)

	// then
	if gc != gear.KeepCurrent {
		t.Errorf("Expected keep gear, got %v", gc)
	}
}

func newEcoDriver(ed externalsystemsfacade.Data) DDrive {
	return NewDDrive(ed)
}

func newExternalData() externalsystemsfacade.Stub {
	return externalsystemsfacade.Stub{}
}
