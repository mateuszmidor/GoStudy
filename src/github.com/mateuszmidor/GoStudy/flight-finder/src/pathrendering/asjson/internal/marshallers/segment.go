package marshallers

import (
	"encoding/json"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/pathrendering/asjson/internal/views"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/segments"
)

// Segment implements segment -> json mashalling
type Segment struct {
	SegmentID segments.ID
	Data      *Data
}

// MarshalJSON implements json.Marshaller for custom marshalling
func (s *Segment) MarshalJSON() ([]byte, error) {
	seg := &s.Data.Segments[s.SegmentID]
	carrier := &s.Data.Carriers[seg.Carrier()]
	airport := &s.Data.Airports[seg.To()]
	nationID := s.Data.Nations.GetByCode(airport.Nation())
	var nation *nations.Nation
	if nationID != nations.NullID {
		nation = &s.Data.Nations[nationID]
	}
	view := views.NewJSONSegmentView(carrier, airport, nation)
	return json.Marshal(view)
}
