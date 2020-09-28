package csv_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading/csv"
)

func TestLoadValidCSVShouldReturnAllAirports(t *testing.T) {
	// given
	var loader csv.AirportLoader
	actualAirports := make(chan dataloading.RawAirport, 1)
	expectedAirports := []dataloading.RawAirport{
		{AirportCode: "GDY", FullName: "Gdynia", Nation: "PL", Latitude: 52.0, Longitude: 18.0},
		{AirportCode: "WAW", FullName: "Warszawa", Nation: "PL", Latitude: 51.5, Longitude: 17.5},
		{AirportCode: "KRK", FullName: "Kraków", Nation: "PL", Latitude: 50.01, Longitude: 19.01},
	}
	// MARKET,LATDEG,LATMIN,LATSEC,LNGDEG,LNGMIN,LNGSEC,LATHEM,LNGHEM,DESCRIPTION
	csv := `
"GDY",52,0,0,18,0,0,"N","E","PL","Gdynia"
"WAW",51,30,0,17,30,0,"N","E","PL","Warszawa"
"KRK",50,0,36,19,0,36,"N","E","PL","Kraków"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualAirports)

	// then
	errorDetails := checkExpectedAirports(expectedAirports, actualAirports)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func TestLoadBrokenCSVShouldReturnOnlyValidAirports(t *testing.T) {
	// given
	var loader csv.AirportLoader
	actualAirports := make(chan dataloading.RawAirport, 1)
	expectedAirports := []dataloading.RawAirport{
		{AirportCode: "KRK", FullName: "Kraków", Nation: "PL", Latitude: 50.01, Longitude: 19.01},
	}
	// MARKET,LATDEG,LATMIN,LATSEC,LNGDEG,LNGMIN,LNGSEC,LATHEM,LNGHEM,NATION,DESCRIPTION
	csv := `
"GDY",52,0,0,18,0,0,"N","E"
"WAW",51,30,0,17,30,0,"N","E","PL", "Warszawa"
"KRK",50,0,36,19,0,36,"N","E","PL","Kraków"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualAirports)

	// then
	errorDetails := checkExpectedAirports(expectedAirports, actualAirports)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func checkExpectedAirports(expected []dataloading.RawAirport, actual chan dataloading.RawAirport) string {
	var result string
	for seg := range actual {
		if index := findAirport(seg, expected); index != -1 {
			removeAirport(index, &expected)
		} else {
			result += fmt.Sprintf("Unexpected Airport loaded: %+v\n", seg)
		}
	}

	if len(expected) != 0 {
		result += fmt.Sprintf("Expected Airports not loaded: %+v", expected)
	}

	return result
}

func findAirport(subject dataloading.RawAirport, list []dataloading.RawAirport) int {
	for i, seg := range list {
		if seg == subject {
			return i
		}
	}
	return -1
}

func removeAirport(index int, list *[]dataloading.RawAirport) {
	l := *list
	l[index] = l[len(l)-1]
	l = l[:len(l)-1]
	*list = l
}
