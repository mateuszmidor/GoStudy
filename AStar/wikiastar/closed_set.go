package wikiastar

type closedNodeSet = map[*Node]bool

type ClosedSet struct {
	nodes closedNodeSet
}

func NewClosedSet() *ClosedSet {
	return &ClosedSet{make(closedNodeSet)}
}

func (s *ClosedSet) Add(node *Node) {
	s.nodes[node] = true
}

func (s *ClosedSet) Contains(node *Node) bool {
	_, exists := s.nodes[node]
	return exists
}
