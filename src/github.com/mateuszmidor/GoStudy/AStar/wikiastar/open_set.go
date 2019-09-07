package wikiastar

type openNodeSet = map[*Node]bool

type OpenSet struct {
	nodes openNodeSet
}

func NewOpenSet() *OpenSet {
	return &OpenSet{make(openNodeSet)}
}

func (s *OpenSet) Add(node *Node) {
	s.nodes[node] = true
}

func (s *OpenSet) Remove(node *Node) {
	delete(s.nodes, node)
}

func (s *OpenSet) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *OpenSet) GetNodeWithLowestFScore(fScore *Score) *Node {
	minCost := fScore.GetInfinity()
	var minCostNode *Node
	for node := range s.nodes {
		cost := fScore.Get(node)
		if cost < minCost {
			minCost = cost
			minCostNode = node
		}
	}
	return minCostNode
}
