package genericastar

type Path = []NodeID

type CameFrom = map[NodeID]NodeID

func reconstructPath(cameFrom CameFrom, current NodeID) Path {
	var ok bool
	path := Path{current}
	for {
		if current, ok = cameFrom[current]; !ok {
			break
		}
		path = append([]NodeID{current}, path...)
	}
	return path
}

func findPath(start NodeID, goal NodeID, neighbors *Neighbors, costs Costs) Path {
	openSet := NewOpenSet()
	openSet.Add(start)
	closedSet := NewClosedSet()

	cameFrom := make(CameFrom)

	gScore := NewScore()
	gScore.Set(start, 0)

	fScore := NewScore()
	fScore.Set(start, costs.H(start, goal))

	for openSet.IsEmpty() == false {
		current := openSet.GetNodeWithLowestFScore(fScore)
		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		openSet.Remove(current)
		closedSet.Add(current)
		for neighbor := range neighbors.GetNeighbors(current) {
			if closedSet.Contains(neighbor) {
				continue
			}

			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentativeGScore is the distance from start to the neighbor through current
			tentativeGScore := gScore.Get(current) + costs.D(current, neighbor)
			if tentativeGScore < gScore.Get(neighbor) {
				cameFrom[neighbor] = current
				gScore.Set(neighbor, tentativeGScore)
				fScore.Set(neighbor, tentativeGScore+costs.H(neighbor, goal))
				openSet.Add(neighbor)
			}
		}
	}

	return nil
}
