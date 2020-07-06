package multipathastar

type ClosedSet struct {
	nodes []bool
}

func NewClosedSet() *ClosedSet {
	return &ClosedSet{make([]bool, 100)}
}

func (s *ClosedSet) Add(node NodeID) {
	if int(node) >= len(s.nodes) {
		newNodes := make([]bool, node*2)
		copy(newNodes, s.nodes)
		s.nodes = newNodes
	}
	s.nodes[node] = true
}

func (s *ClosedSet) Remove(node NodeID) {
	if int(node) >= len(s.nodes) {
		return
	}
	s.nodes[node] = false
}

func (s *ClosedSet) Contains(node NodeID) bool {
	return int(node) < len(s.nodes) && s.nodes[node]
}
