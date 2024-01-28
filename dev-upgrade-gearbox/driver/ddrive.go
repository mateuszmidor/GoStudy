package gearboxdriver

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/aggressiveness"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/characteristics"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/drivingmode"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/externalsystemsfacade"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/gearadjustment"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/mdynamic"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/events"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gas"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gear"
)

// DDrive exhibits gear adjustment functionality for D-Drive mode
type DDrive struct {
	mdynamic            mdynamic.Mode
	aggressivenessLevel aggressiveness.Level
	drivingMode         drivingmode.Mode
	externalData        externalsystemsfacade.Data
}

// NewDDrive is constructor
func NewDDrive(ed externalsystemsfacade.Data) DDrive {
	d := DDrive{
		externalData: ed,
	}
	d.SetDrivingModeEco()
	d.SetAggressivenessLevel1()
	return d
}

// HandleGas returns how current gear should change to accomodate current situation + events
func (d *DDrive) HandleGas(gas gas.Value) (totalChange gear.Change, eventList events.Events) {
	totalChange = gear.KeepCurrent

	// Check for manual gear change active
	if d.activeManualGearChangePreventsAutomaticGearChange() {
		return
	}

	// Check for mdynamic mode in uslizg
	if d.mDynamicPreventsGearChange() {
		return
	}

	// Regular driving - adjust for Optimal RPM
	changeForOptimalRPM := d.adjustGearForOptimalRPM()
	totalChange = totalChange.Add(changeForOptimalRPM)

	// Descending with trailor - adjust for Engine breaking
	changeForEngineBreaking := d.adjustGearForEngineBreaking(gas)
	totalChange = totalChange.Add(changeForEngineBreaking)

	// Overtaking - adjust for KickDown
	changeForKickDown := d.adjustGearForKickDown(gas)
	totalChange = totalChange.Add(changeForKickDown)

	// gear change emits events
	eventList = d.emitEvents(totalChange)

	return
}

// SetMDynamic setter
func (d *DDrive) SetMDynamic(enabled bool) {
	if enabled {
		d.mdynamic.Enable()
	}
	// disabling not needed for now
}

// SetAggressivenessLevel1 setter
func (d *DDrive) SetAggressivenessLevel1() {
	d.aggressivenessLevel = aggressiveness.NewLevel1()
}

// SetAggressivenessLevel2 setter
func (d *DDrive) SetAggressivenessLevel2() {
	panic("not implemented")
}

// SetAggressivenessLevel3 setter
func (d *DDrive) SetAggressivenessLevel3() {
	d.aggressivenessLevel = aggressiveness.NewLevel3()
}

// SetDrivingModeEco setter
func (d *DDrive) SetDrivingModeEco() {
	minRPM, maxRPM := characteristics.GetRPMForSpeedingUpEco()
	d.drivingMode = drivingmode.NewEco(minRPM, maxRPM)
}

// SetDrivingModeSport setter
func (d *DDrive) SetDrivingModeSport() {
	minRPM, maxRPM := characteristics.GetRPMForSpeedingUpSport()
	kickdown1 := gas.New(characteristics.GetGasKickDown1Sport())
	kickdown2 := gas.New(characteristics.GetGasKickDown2Sport())
	d.drivingMode = drivingmode.NewSport(minRPM, maxRPM, kickdown1, kickdown2)
}

// SetDrivingModeComfort setter
func (d *DDrive) SetDrivingModeComfort() {
	panic("not implemented")
	// minRPM, maxRPM := characteristics.GetRPMForSpeedingUpComfort()
	// d.drivingMode = drivingmode.NewComfort(minRPM, maxRPM)
}

func (d *DDrive) activeManualGearChangePreventsAutomaticGearChange() bool {
	return d.externalData.IsManualGearChangeActive()
}

func (d *DDrive) mDynamicPreventsGearChange() bool {
	return !d.mdynamic.IsGearChangeAllowed(d.externalData.GetAngularSpeed())
}

func (d *DDrive) adjustGearForOptimalRPM() gear.Change {
	minRPM, maxRPM := d.drivingMode.GetOptimalRPM(d.aggressivenessLevel)
	return gearadjustment.AdjustForOptimalRPM(d.externalData.GetCurrentRPM(), minRPM, maxRPM)
}

func (d *DDrive) adjustGearForEngineBreaking(gas gas.Value) gear.Change {
	if !gas.IsZero() {
		return gear.KeepCurrent
	}
	hasTrailor := d.externalData.GetTrailorAttached()
	isGoingDownTheSlope := d.externalData.GetDrivingDownTheSlope()
	return gearadjustment.AdjustGearForEngineBreaking(hasTrailor, isGoingDownTheSlope)
}

func (d *DDrive) adjustGearForKickDown(gas gas.Value) gear.Change {
	return gearadjustment.AdjustForKickDown(d.drivingMode, gas)
}

func (d *DDrive) emitEvents(gc gear.Change) events.Events {
	switch gc {
	case gear.GearDown:
		return events.Events{events.GearDown}
	case gear.DoubleGearDown:
		return events.Events{events.DoubleGearDown}
	case gear.GearUp:
		return events.Events{events.GearUp}
	default:
		return events.Events{}
	}
}
