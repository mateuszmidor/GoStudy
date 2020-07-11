package connections

import (
	"airports"
	"pathfinding"
	"segments"
)

// Adapter is pathfinding.Connections adapter for segments
type Adapter struct {
	segments segments.Segments // assumption: segments are sorted ascending by from, to
}

// NewAdapter is constructor
func NewAdapter(segments segments.Segments) Adapter {
	return Adapter{segments}
}

// GetDestinationNode implements Connections interface
func (a *Adapter) GetDestinationNode(connection pathfinding.ConnectionID) pathfinding.NodeID {
	return pathfinding.NodeID(a.segments[connection].To())
}

// GetOutgoingConnections implements Connections interface
func (a *Adapter) GetOutgoingConnections(node pathfinding.NodeID) (first, last pathfinding.ConnectionID) {
	var finder SegmentRangeFinder
	f, l := finder.ByFromAirport(a.segments, airports.ID(node))
	return pathfinding.ConnectionID(f), pathfinding.ConnectionID(l)
}
