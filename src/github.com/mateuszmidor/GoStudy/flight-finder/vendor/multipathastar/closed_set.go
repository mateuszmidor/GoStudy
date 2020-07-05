package multipathastar

type closedNodeSet = map[NodeID]bool

type ClosedSet struct {
	nodes closedNodeSet
}

func NewClosedSet() *ClosedSet {
	return &ClosedSet{make(closedNodeSet)}
}

func (s *ClosedSet) Add(node NodeID) {
	s.nodes[node] = true
}

func (s *ClosedSet) Remove(node NodeID) {
	delete(s.nodes, node)
}

func (s *ClosedSet) Contains(node NodeID) bool {
	_, exists := s.nodes[node]
	return exists
}
