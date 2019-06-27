package cluster

import "math/rand"

func RandomStation(stationList []string) uint32 {
	// rand.Intn(0) results in painc()
	if numStations := len(stationList); numStations == 0 {
		return 0
	} else {
		return uint32(rand.Intn(numStations) + 3) // + 3 to generate some invalid tune commands
	}
}