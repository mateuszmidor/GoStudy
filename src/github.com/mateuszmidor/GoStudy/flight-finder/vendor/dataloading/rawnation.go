package dataloading

// RawNation is nation in its denormalized form
type RawNation struct {
	Code     string // eg. "PL"
	Iso      string // eg. "POL"
	Currency string // eg. "PLN"
	Name     string // eg. "POLAND"
}

// NewRawNation is constructor
func NewRawNation(code, iso, currency, name string) RawNation {
	return RawNation{
		Code:     code,
		Iso:      iso,
		Currency: currency,
		Name:     name,
	}
}
