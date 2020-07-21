package asjson_test

import (
	"airports"
	"bytes"
	"carriers"
	"encoding/json"
	"geo"
	"pathfinding"
	"pathrendering/asjson"
	"segments"
	"testing"
)

type jsonpath struct {
	FromAirport jsonairport   `json:"from_airport"`
	Segments    []jsonsegment `json:"segments"`
}

type jsonsegment struct {
	Carrier   jsoncarrier `json:"carrier"`
	ToAirport jsonairport `json:"to_airport"`
}

type jsonairport struct {
	Code      string  `json:"code"`
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type jsoncarrier struct {
	Code string `json:"code"`
}

func TestPathRendererTurnsValidPathIntoValidPathJson(t *testing.T) {
	// given
	// notice: airports sorted by code
	airports := airports.Airports{
		airports.NewAirport("GDN", "GDANSK", "PL", geo.Longitude(51), geo.Latitude(21)),
		airports.NewAirport("KRK", "KRAKOW", "PL", geo.Longitude(49), geo.Latitude(19)),
		airports.NewAirport("WAW", "WARSZAWA", "PL", geo.Longitude(50), geo.Latitude(20)),
	}
	// notice: carriers sorted by code
	carriers := carriers.Carriers{
		carriers.NewCarrier("FR"),
		carriers.NewCarrier("LO"),
	}
	// notice: segments sorted by from airport
	segments := segments.Segments{
		segments.NewSegment(1, 2, 1), // connectionID=0 : KRK-WAW
		segments.NewSegment(2, 0, 0), // connectionID=1 : WAW-GDN
	}
	path := pathfinding.Path{
		pathfinding.ConnectionID(0),
		pathfinding.ConnectionID(1),
	}

	// KRK-(LO)-WAW-(FR)-GDN
	expected := jsonpath{
		FromAirport: jsonairport{
			Code:      "KRK",
			Longitude: 49.0,
			Latitude:  19.0,
		},
		Segments: []jsonsegment{
			{
				Carrier: jsoncarrier{
					Code: "LO",
				},
				ToAirport: jsonairport{
					Code:      "WAW",
					Longitude: 50.0,
					Latitude:  20.0,
				},
			},
			{
				Carrier: jsoncarrier{
					Code: "FR",
				},
				ToAirport: jsonairport{
					Code:      "GDN",
					Longitude: 51.0,
					Latitude:  21.0,
				},
			},
		},
	}
	buf := bytes.NewBuffer([]byte{})
	renderer := asjson.NewPathRenderer(airports, carriers, segments)

	// when
	renderer.Render(buf, []pathfinding.Path{path})

	// then
	var actualPaths []jsonpath
	json.NewDecoder(buf).Decode(&actualPaths)
	if len(actualPaths) != 1 {
		t.Fatalf("For single input path there should be single path outputted in json, got %d", len(actualPaths))
	}

	actual := actualPaths[0]
	if actual.FromAirport != expected.FromAirport {
		t.Errorf("For path %v the expected from airport is %+v, got %+v", path, expected.FromAirport, actual.FromAirport)
	}

	if len(actual.Segments) != len(expected.Segments) {
		t.Fatalf("For path %v the expected num of segments is %d, got %d", path, len(expected.Segments), len(actual.Segments))
	}

	if actual.Segments[0] != expected.Segments[0] {
		t.Errorf("For path %v the expected first segment is %+v, got %+v", path, expected.Segments[0], actual.Segments[0])
	}

	if actual.Segments[1] != expected.Segments[1] {
		t.Errorf("For path %v the expected second segment is %+v, got %+v", path, expected.Segments[1], actual.Segments[1])
	}
}

func TestPathRendererTurnsEmptyPathsIntoEmptyJSON(t *testing.T) {
	// given
	emptyJSONArray := "[]\n"
	airports := airports.Airports{}
	carriers := carriers.Carriers{}
	segments := segments.Segments{}
	buf := bytes.NewBuffer([]byte{})
	renderer := asjson.NewPathRenderer(airports, carriers, segments)

	// when
	renderer.Render(buf, []pathfinding.Path{})

	// then
	if buf.String() != emptyJSONArray {
		t.Errorf("For empty paths, there should be empty json array, got %q", buf.String())
	}
}
