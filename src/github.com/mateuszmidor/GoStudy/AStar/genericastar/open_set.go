package genericastar

type openNodeSet = map[NodeID]bool

type OpenSet struct {
	nodes openNodeSet
}

func NewOpenSet() *OpenSet {
	return &OpenSet{make(openNodeSet)}
}

func (s *OpenSet) Add(node NodeID) {
	s.nodes[node] = true
}

func (s *OpenSet) Remove(node NodeID) {
	delete(s.nodes, node)
}

func (s *OpenSet) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *OpenSet) GetNodeWithLowestFScore(fScore *Score) NodeID {
	minCost := fScore.GetInfinity()
	var minCostNode NodeID
	for node := range s.nodes {
		cost := fScore.Get(node)
		if cost < minCost {
			minCost = cost
			minCostNode = node
		}
	}
	return minCostNode
}
