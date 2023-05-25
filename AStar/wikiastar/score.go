package wikiastar

import "math"

type ScoreType = float64
type nodeToScoreMap = map[*Node]ScoreType

type Score struct {
	scores nodeToScoreMap
}

func NewScore() *Score {
	return &Score{make(nodeToScoreMap)}
}

func (s *Score) GetInfinity() ScoreType {
	return math.MaxFloat64
}

func (s *Score) Get(n *Node) ScoreType {
	if val, ok := s.scores[n]; ok {
		return val
	}
	return s.GetInfinity()
}

func (s *Score) Set(n *Node, score ScoreType) {
	s.scores[n] = score
}
