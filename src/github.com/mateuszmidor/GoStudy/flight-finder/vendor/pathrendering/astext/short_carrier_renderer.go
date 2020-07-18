package astext

import (
	"carriers"
)

// ShortCarrierRenderer renders short string representation of carrierID
type ShortCarrierRenderer struct {
	carriers carriers.Carriers
}

// NewShortCarrierRenderer is constructor
func NewShortCarrierRenderer(carriers carriers.Carriers) *ShortCarrierRenderer {
	return &ShortCarrierRenderer{carriers}
}

// Render creates string representation of carrier
func (r *ShortCarrierRenderer) Render(id carriers.ID) string {
	return r.carriers[id].Code()
}
