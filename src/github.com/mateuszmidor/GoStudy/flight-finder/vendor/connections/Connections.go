package connections

import (
	"airport"
	"pathfinding"
	"segment"
)

// Connections is flight connection network, implements pathfinding.Connections
type Connections struct {
	segments segment.Segments // assumption: segments are sorted ascending by from, to
}

// NewConnections is constructor
func NewConnections(segments segment.Segments) Connections {
	return Connections{segments}
}

// GetDestinationNode implements Connections interface
func (n *Connections) GetDestinationNode(connection pathfinding.ConnectionID) pathfinding.NodeID {
	return pathfinding.NodeID(n.segments[connection].To())
}

// GetOutgoingConnections implements Connections interface
func (n *Connections) GetOutgoingConnections(node pathfinding.NodeID) (first, last pathfinding.ConnectionID) {
	var finder SegmentRangeFinder
	f, l := finder.ByFromAirport(n.segments, airport.ID(node))
	return pathfinding.ConnectionID(f), pathfinding.ConnectionID(l)
}
