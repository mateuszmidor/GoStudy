package uslizg

import "github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/types"

const uślizgThreshold types.AngularSpeed = 60

// AngularSpeedForNoUślizg represents AngularSpeed for no-uslizg
const AngularSpeedForNoUślizg types.AngularSpeed = uślizgThreshold - 10

// AngularSpeedForUślizg represents AngularSpeed for uslizg
const AngularSpeedForUślizg types.AngularSpeed = uślizgThreshold + 10

// IsUślizg checks if "as" value reached uslizg threshold
func IsUślizg(as types.AngularSpeed) bool {
	return as >= uślizgThreshold
}
