package marshallers

import (
	"airports"
	"encoding/json"
	"pathrendering/asjson/internal/views"
)

// Airport implements airport -> json mashalling
type Airport struct {
	AirportID airports.ID
	Data      *Data
}

// MarshalJSON implements json.Marshaller for custom marshalling
func (a *Airport) MarshalJSON() ([]byte, error) {
	view := views.NewJSONAirportView(&a.Data.Airports[a.AirportID])
	return json.Marshal(view)
}
