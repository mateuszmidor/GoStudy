package wikiastar

import "math"

func distance(from *Node, to *Node) float64 {
	dx := float64(to.X - from.X)
	dy := float64(to.Y - from.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// cost of edge from-to
func d(from *Node, to *Node) float64 {
	return distance(from, to)
}

// heuristic cost of path current-goal
func h(current *Node, goal *Node) float64 {
	return distance(current, goal)
}

type PathFinder struct {
}

type CameFrom = map[*Node]*Node

func reconstructPath(cameFrom map[*Node]*Node, current *Node) string {
	var ok bool
	path := current.Name
	for {
		if current, ok = cameFrom[current]; !ok {
			break
		}
		path = current.Name + "-" + path
	}
	return path
}

func findPath(start *Node, goal *Node, graph *Graph) string {
	openSet := NewOpenSet()
	openSet.Add(start)
	closedSet := NewClosedSet()

	cameFrom := make(CameFrom)

	gScore := NewScore()
	gScore.Set(start, 0)

	fScore := NewScore()
	fScore.Set(start, h(start, goal))

	for openSet.IsEmpty() == false {
		current := openSet.GetNodeWithLowestFScore(fScore)
		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		openSet.Remove(current)
		closedSet.Add(current)
		for neighbor := range graph.GetNeighbors(current) {
			if closedSet.Contains(neighbor) {
				continue
			}

			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentativeGScore is the distance from start to the neighbor through current
			tentativeGScore := gScore.Get(current) + d(current, neighbor)
			if tentativeGScore < gScore.Get(neighbor) {
				cameFrom[neighbor] = current
				gScore.Set(neighbor, tentativeGScore)
				fScore.Set(neighbor, tentativeGScore+h(neighbor, goal))
				openSet.Add(neighbor)
			}
		}
	}

	return "[path not found]"
}
