package connections

import (
	"airport"
	"multipathastar"
	"segment"
)

// Network is flight connection network, implements multipathastar.Connections
type Network struct {
	airports airport.Airports
	segments segment.Segments // assumption: segments are sorted ascending by from, to
}

// NewNetwork is constructor
func NewNetwork(airports airport.Airports, segments segment.Segments) Network {
	return Network{airports, segments}
}

// GetDestinationNode implements Connections interface
func (n *Network) GetDestinationNode(connection multipathastar.ConnectionID) multipathastar.NodeID {
	return multipathastar.NodeID(n.segments[connection].To())
}

// GetOutgoingConnections implements Connections interface
func (n *Network) GetOutgoingConnections(node multipathastar.NodeID) (first, last multipathastar.ConnectionID) {
	var finder SegmentRangeFinder
	f, l := finder.ByFromAirport(n.segments, airport.AirportID(node))
	return multipathastar.ConnectionID(f), multipathastar.ConnectionID(l)
}

// GetAirport returns airport of given id
func (n *Network) GetAirport(id airport.AirportID) *airport.Airport {
	return &n.airports[id]
}

// GetSegment returns segment of given id
func (n *Network) GetSegment(id segment.ID) *segment.Segment {
	return &n.segments[id]
}
