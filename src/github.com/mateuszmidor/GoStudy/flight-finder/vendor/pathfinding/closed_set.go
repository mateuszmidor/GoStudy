package pathfinding

type ClosedSet struct {
	nodes []bool
	size  int
}

func NewClosedSet() *ClosedSet {
	size := 100
	nodes := make([]bool, size)
	return &ClosedSet{nodes, size}
}

func (s *ClosedSet) Add(node NodeID) {
	if int(node) >= s.size {
		s.size = int(node) * 2
		newNodes := make([]bool, s.size)
		copy(newNodes, s.nodes)
		s.nodes = newNodes
	}
	s.nodes[node] = true
}

func (s *ClosedSet) Remove(node NodeID) {
	if int(node) >= s.size {
		return
	}
	s.nodes[node] = false
}

func (s *ClosedSet) Contains(node NodeID) bool {
	if int(node) >= s.size {
		return false
	}
	return s.nodes[node]
}
