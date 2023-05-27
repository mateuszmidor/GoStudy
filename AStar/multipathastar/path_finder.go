package multipathastar

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

// find all possible paths, not only the least costly
func findPaths(start NodeID, goal NodeID, neighbors *Neighbors, costs Costs) (result []Path) {
	closedSet := NewClosedSet()
	cameFrom := make(CameFrom)
	var recursive func(current NodeID)

	recursive = func(current NodeID) {
		// check if path found
		if current == goal {
			result = append(result, reconstructPath(cameFrom, current))
			return
		}

		// mark current as visited
		closedSet.Add(current)

		// check all neighbors of current
		for neighbor := range neighbors.GetNeighbors(current) {
			// avoid cycles
			if closedSet.Contains(neighbor) {
				continue
			}

			// process neighbor
			cameFrom[neighbor] = current
			recursive(neighbor)
			closedSet.Remove(neighbor)
		}
	}

	recursive(start)
	return result
}
