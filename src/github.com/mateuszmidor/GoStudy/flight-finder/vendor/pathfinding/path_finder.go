package pathfinding

// FindPaths find all possible paths, not only the least costly
func FindPaths(start NodeID, goal NodeID, connections Connections) (result []Path) {
	closedSet := NewClosedSet()
	currPath := make(Path, 3)

	var recursive func(current NodeID, pathLength int)
	recursive = func(current NodeID, pathLength int) {
		if pathLength >= len(currPath) {
			return
			//panic("Max path length exceeded")
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
		closedSet.Add(current)

		// check all outgoing connections of current
		first, last := connections.GetOutgoingConnections(current)
		for connID := first; connID < last; connID++ {
			dest := connections.GetDestinationNode(connID)

			// avoid cycles
			if closedSet.Contains(dest) {
				continue
			}

			// process outgoing connection
			currPath[pathLength] = connID
			recursive(dest, pathLength+1)

			closedSet.Remove(dest) // dest added by "recursive"
		}

	}

	recursive(start, 0)
	return result
}
