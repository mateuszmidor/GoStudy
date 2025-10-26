package graph

type edges map[string]map[string]bool

// Graph is unidirectional graph
type Graph struct {
	edges edges
}

// NewGraph creates new Graph
func NewGraph() Graph {
	return Graph{make(edges)}
}

// HasEdge returns if graph contains edge from a to b
func (g Graph) HasEdge(a, b string) bool {
	return g.edges[a][b]
}

// AddEdge adds edge a-b to the graph
func (g Graph) AddEdge(a, b string) {
	edges := g.edges[a]
	if edges == nil {
		edges = make(map[string]bool)
		g.edges[a] = edges
	}
	edges[b] = true
}
