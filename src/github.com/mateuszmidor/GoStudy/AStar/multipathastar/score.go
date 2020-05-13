package multipathastar

import "math"

type ScoreType = float64
type nodeToScoreMap = map[NodeID]ScoreType

type Score struct {
	scores nodeToScoreMap
}

func NewScore() *Score {
	return &Score{make(nodeToScoreMap)}
}

func (s *Score) GetInfinity() ScoreType {
	return math.MaxFloat64
}

func (s *Score) Get(n NodeID) ScoreType {
	if val, ok := s.scores[n]; ok {
		return val
	}
	return s.GetInfinity()
}

func (s *Score) Set(n NodeID, score ScoreType) {
	s.scores[n] = score
}
