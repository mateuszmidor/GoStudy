package marshallers

import (
	"airports"
	"encoding/json"
	"nations"
	"pathrendering/asjson/internal/views"
)

// Airport implements airport -> json mashalling
type Airport struct {
	AirportID airports.ID
	Data      *Data
}

// MarshalJSON implements json.Marshaller for custom marshalling
func (a *Airport) MarshalJSON() ([]byte, error) {
	airport := &a.Data.Airports[a.AirportID]
	nationID := a.Data.Nations.GetByCode(airport.Nation())
	var nation *nations.Nation
	if nationID != nations.NullID {
		nation = &a.Data.Nations[nationID]
	}
	view := views.NewJSONAirportView(airport, nation)
	return json.Marshal(view)
}
