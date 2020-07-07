package connections

import (
	"airport"
	"pathfinding"
	"segment"
)

// Network is flight connection network, implements pathfinding.Connections
type Network struct {
	airports airport.Airports
	segments segment.Segments // assumption: segments are sorted ascending by from, to
}

// NewNetwork is constructor
func NewNetwork(airports airport.Airports, segments segment.Segments) Network {
	return Network{airports, segments}
}

// GetDestinationNode implements Connections interface
func (n *Network) GetDestinationNode(connection pathfinding.ConnectionID) pathfinding.NodeID {
	return pathfinding.NodeID(n.segments[connection].To())
}

// GetOutgoingConnections implements Connections interface
func (n *Network) GetOutgoingConnections(node pathfinding.NodeID) (first, last pathfinding.ConnectionID) {
	var finder SegmentRangeFinder
	f, l := finder.ByFromAirport(n.segments, airport.AirportID(node))
	return pathfinding.ConnectionID(f), pathfinding.ConnectionID(l)
}

// GetAirport returns airport of given id
func (n *Network) GetAirport(id airport.AirportID) *airport.Airport {
	return &n.airports[id]
}

// GetSegment returns segment of given id
func (n *Network) GetSegment(id segment.ID) *segment.Segment {
	return &n.segments[id]
}
