package segment

// RawSegment is segment in its denormalized form
type RawSegment struct {
	FromAirportCode, ToAirportCode string // eg. "KTW"
	CarrierCode                    string // eg. "LO"
}
