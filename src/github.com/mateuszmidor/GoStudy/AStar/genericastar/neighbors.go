package genericastar

type NeighborSet = map[NodeID]bool
type neighborMap = map[NodeID]NeighborSet

type Neighbors struct {
	neighbors neighborMap
}

func NewNeighbors() *Neighbors {
	return &Neighbors{make(neighborMap)}
}

func (n *Neighbors) GetNeighbors(node NodeID) NeighborSet {
	if neighbors, ok := n.neighbors[node]; ok {
		return neighbors
	}
	return NeighborSet{}
}

func (n *Neighbors) getOrCreateNeighbors(node NodeID) NeighborSet {
	// if exists - then return
	if neighbors, ok := n.neighbors[node]; ok {
		return neighbors
	}

	// otherwise insert and return
	n.neighbors[node] = make(NeighborSet)
	return n.neighbors[node]
}

func (n *Neighbors) Connect(from NodeID, to NodeID) {
	neighbors := n.getOrCreateNeighbors(from)
	neighbors[to] = true
}
