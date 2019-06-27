package hardware

import "math/rand"

func RandomStationList() []string {
	stationList:= []string{"40'", "50'", "60'", "70'", "80'", "90'", "Elvis Radio", "1st Wave", "Velvet"}
	result:= []string{}
	numStations := rand.Intn(len(stationList))
	for i:= 0; i < numStations; i++ {
		// pick random station
		index := rand.Intn(len(stationList))

		// add station to result list
		result = append(result, stationList[index])

		// remove station from the collection
		stationList = append(stationList[:index], stationList[index+1:]...)
	}
	return result
}