package asjson

import "carriers"

// CarrierView is json view of carrieres.Carrier
type CarrierView struct {
	Code string `json:"code"`
}

// NewJSONCarrierView is constructor
func NewJSONCarrierView(c *carriers.Carrier) *CarrierView {
	return &CarrierView{
		Code: c.Code(),
	}
}
