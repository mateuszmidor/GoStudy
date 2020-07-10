package connections

import (
	"airport"
	"pathfinding"
	"segment"
)

// Network is flight connection network, implements pathfinding.Connections
type Network struct {
	segments segment.Segments // assumption: segments are sorted ascending by from, to
}

// NewNetwork is constructor
func NewNetwork(segments segment.Segments) Network {
	return Network{segments}
}

// GetDestinationNode implements Connections interface
func (n *Network) GetDestinationNode(connection pathfinding.ConnectionID) pathfinding.NodeID {
	return pathfinding.NodeID(n.segments[connection].To())
}

// GetOutgoingConnections implements Connections interface
func (n *Network) GetOutgoingConnections(node pathfinding.NodeID) (first, last pathfinding.ConnectionID) {
	var finder SegmentRangeFinder
	f, l := finder.ByFromAirport(n.segments, airport.ID(node))
	return pathfinding.ConnectionID(f), pathfinding.ConnectionID(l)
}
