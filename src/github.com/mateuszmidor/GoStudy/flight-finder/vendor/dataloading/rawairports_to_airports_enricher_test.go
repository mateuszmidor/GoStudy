package dataloading_test

import (
	"airports"
	"dataloading"
	"geo"
	"testing"
)

func TestEnricherAddsExpectedInformationToAirports(t *testing.T) {
	// given
	rawAirports := make(chan dataloading.RawAirport, 3)
	rawAirports <- dataloading.NewRawAirport("WAW", "Warszawa", "PL", geo.Longitude(20.0), geo.Latitude(51.0))
	rawAirports <- dataloading.NewRawAirport("GDN", "Gdańsk", "PL", geo.Longitude(21.0), geo.Latitude(52.0))
	rawAirports <- dataloading.NewRawAirport("KRK", "Kraków", "PL", geo.Longitude(19.0), geo.Latitude(50.0))
	close(rawAirports)
	// airports need be sorted by code for binary search purpose
	actualAirports := airports.Airports{
		airports.NewAirportCodeOnly("GDN"),
		airports.NewAirportCodeOnly("KRK"),
		airports.NewAirportCodeOnly("WAW"),
	}
	expectedAirports := airports.Airports{
		airports.NewAirport("GDN", "Gdańsk", "PL", geo.Longitude(21.0), geo.Latitude(52.0)),
		airports.NewAirport("KRK", "Kraków", "PL", geo.Longitude(19.0), geo.Latitude(50.0)),
		airports.NewAirport("WAW", "Warszawa", "PL", geo.Longitude(20.0), geo.Latitude(51.0)),
	}

	// when
	dataloading.EnrichAirports(actualAirports, rawAirports)

	// then
	for i, actual := range actualAirports {
		if actual != expectedAirports[i] {
			t.Errorf("For index %d expected airport %+v, got %+v", i, expectedAirports[i], actual)
		}
	}
}
