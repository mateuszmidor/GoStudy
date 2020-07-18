package marshallers

import (
	"encoding/json"
	"pathrendering/asjson/internal/views"
	"segments"
)

// Segment implements segment -> json mashalling
type Segment struct {
	SegmentID segments.ID
	Data      *Data
}

// MarshalJSON implements json.Marshaller for custom marshalling
func (s *Segment) MarshalJSON() ([]byte, error) {
	seg := &s.Data.Segments[s.SegmentID]
	view := views.NewJSONSegmentView(&s.Data.Carriers[seg.Carrier()], &s.Data.Airports[seg.To()])
	return json.Marshal(view)
}
