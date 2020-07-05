package connections

import (
	"airport"
	"multipathastar"
	"segment"
)

// Network is flight connection network, implements multipathastar.Connections
type Network struct {
	airports airport.Airports
	segments segment.Segments
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
	f, l := equalRange(n.segments, airport.AirportID(node))
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

func equalRange(segments segment.Segments, id airport.AirportID) (first, last segment.ID) {
	first = -1
	last = -2 // so for i := first; i <= last; i++ doesnt execute when no elements

	for i := 0; i < len(segments); i++ {
		if segments[i].From() == id {
			first = segment.ID(i)
			break
		}
	}

	for i := len(segments) - 1; i >= 0; i-- {
		if segments[i].From() == id {
			last = segment.ID(i)
			break
		}
	}

	return
}
