package asjson

import (
	"pathfinding"
	"pathrendering/asjson/internal/marshallers"
	"segments"
)

func buildMarshallerForPaths(paths []pathfinding.Path, d *marshallers.Data) []*marshallers.Path {
	result := make([]*marshallers.Path, 0, len(paths))
	for i := range paths {
		if marshaller := buildMarshallerForPath(paths[i], d); marshaller != nil {
			result = append(result, marshaller)
		}
	}

	return result
}

func buildMarshallerForPath(path pathfinding.Path, d *marshallers.Data) *marshallers.Path {
	if len(path) == 0 {
		return nil
	}

	s0 := &d.Segments[path[0]]
	result := marshallers.Path{
		FromAirport: marshallers.Airport{AirportID: s0.From(), Data: d},
		Segments:    make([]marshallers.Segment, len(path)),
	}

	for i, sID := range path {
		result.Segments[i] = marshallers.Segment{SegmentID: segments.ID(sID), Data: d}
	}

	return &result
}
