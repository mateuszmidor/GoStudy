package uślizg

import "driver/types"

const uślizgThreshold types.AngularSpeed = 60

// AngularSpeedForNoUślizg represents AngularSpeed for no-uślizg
const AngularSpeedForNoUślizg types.AngularSpeed = uślizgThreshold - 10

// AngularSpeedForUślizg represents AngularSpeed for uślizg
const AngularSpeedForUślizg types.AngularSpeed = uślizgThreshold + 10

// IsUślizg checks if "as" value reached uślizg threshold
func IsUślizg(as types.AngularSpeed) bool {
	return as >= uślizgThreshold
}
