package application

import "math/rand"

func RandomStation(stationList []string) uint32 {
	numStations := len(stationList)
	return uint32(rand.Intn(numStations + 1)) // + 1 to generate some invalid tune commands
}
