package segment

// RawSegment is segment in its denormalized form
type RawSegment struct {
	FromAirportCode, ToAirportCode string // eg. "KTW"
	CarrierCode                    string // eg. "LO"
}

func NewRawSegment(from, to, carrier string) RawSegment {
	return RawSegment{FromAirportCode: from, ToAirportCode: to, CarrierCode: carrier}
}
