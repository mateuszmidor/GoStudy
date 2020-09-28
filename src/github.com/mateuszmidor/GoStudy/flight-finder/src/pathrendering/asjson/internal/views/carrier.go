package views

import "github.com/mateuszmidor/GoStudy/flight-finder/src/carriers"

// Carrier is json view of carrieres.Carrier
type Carrier struct {
	Code string `json:"code"`
}

// NewJSONCarrierView is constructor
func NewJSONCarrierView(c *carriers.Carrier) *Carrier {
	return &Carrier{
		Code: c.Code(),
	}
}
