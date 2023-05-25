package wikiastar

type NeighborSet = map[*Node]bool
type NeighborMap = map[*Node]NeighborSet

type Graph struct {
	neighbors NeighborMap
}

func NewGraph() *Graph {
	return &Graph{make(NeighborMap)}
}

func (g *Graph) GetNeighbors(node *Node) NeighborSet {
	if neighbors, ok := g.neighbors[node]; ok {
		return neighbors
	}
	return NeighborSet{}
}

func (g *Graph) getOrCreateNeighbors(node *Node) NeighborSet {
	// if exists - then return
	if _, ok := g.neighbors[node]; ok {
		return g.neighbors[node]
	}

	// otherwise insert and return
	g.neighbors[node] = make(NeighborSet)
	return g.neighbors[node]
}

func (g *Graph) Connect(from *Node, to *Node) {
	neighbors := g.getOrCreateNeighbors(from)
	neighbors[to] = true
}
