package multipathastar

// Connections represents a network of interconnected nodes for the sake of path finding
type Connections interface {
	GetOutgoingConnections(node NodeID) (first ConnectionID, last ConnectionID)
	GetDestinationNode(connection ConnectionID) NodeID
}
