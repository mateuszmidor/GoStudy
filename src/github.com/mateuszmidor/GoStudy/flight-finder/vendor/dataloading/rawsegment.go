package dataloading

// RawSegment is segment in its denormalized form
type RawSegment struct {
	FromAirportCode string // eg. "GDN"
	ToAirportCode   string // eg. "KRK"
	CarrierCode     string // eg. "LO"
}

// NewRawSegment is constructor
func NewRawSegment(from, to, carrier string) RawSegment {
	return RawSegment{
		FromAirportCode: from,
		ToAirportCode:   to,
		CarrierCode:     carrier,
	}
}
