package pathfinding

// FindPaths find all possible paths, not only the least costly
func FindPaths(start NodeID, goal NodeID, connections Connections) (result []Path) {
	const maxNodeID = 1000000 // assumption: nodeID < 1'000'000; for airports should be enough
	var alreadyVisited [maxNodeID]bool
	currPath := make(Path, 3)

	var recursive func(current NodeID, pathLength int)
	recursive = func(current NodeID, pathLength int) {
		if pathLength >= len(currPath) {
			return
		}

		if len(result) >= 1000 {
			return
		}

		// check if path found
		if current == goal {
			newPath := make(Path, pathLength)
			copy(newPath, currPath)
			result = append(result, newPath)
			return
		}

		// mark current as visited. For cycle detection
		alreadyVisited[current] = true

		// check all outgoing connections of current
		first, last := connections.GetOutgoingConnections(current)
		for connID := first; connID < last; connID++ {
			dest := connections.GetDestinationNode(connID)

			// avoid cycles
			if alreadyVisited[dest] {
				continue
			}

			// process outgoing connection
			currPath[pathLength] = connID
			recursive(dest, pathLength+1)

			alreadyVisited[dest] = false // dest added by "recursive"
		}

	}

	recursive(start, 0)
	return result
}
