package marshallers

// Path implements path -> json mashalling
// it represents the path in json as origin airport followed by sequence of (carrier, destination airport) pairs, eg for KRK-WAW-GDN:
// KRK
//   LO -> WAW
//   FR -> GDN
type Path struct {
	FromAirport Airport   `json:"from_airport"`
	Segments    []Segment `json:"segments"`
}
