package problem1319

// Thorem: for n points, we need at least n-1 lines to
// connected n.
//
func makeConnected(n int, connections [][]int) int {

	if len(connections) < n-1 {
		return -1
	}

	// Create a graph by 2D table.
	// g[i] is the neighbors of i.
	// O(E)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}

	for _, c := range connections {
		g[c[0]] = append(g[c[0]], c[1])
		g[c[1]] = append(g[c[1]], c[0])
	}

	// Records the seen nodes.
	// Every node need only be visited once!
	seen := make([]int, n)
	// The current connected custers' count.
	count := 0

	for i := 0; i < n; i++ {
		// If node i has been visited,
		// means i is not a new node needed to be connected.
		if seen[i] != 0 {
			continue
		} else {
			count++
			// Mark all node i's connected node.
			dfs(i, g, seen)
		}
	}
	return count - 1
}

// deep first search to mark all can reached node from cur node.
// From the current node, to check its neighbors,
func dfs(cur int, g [][]int, seen []int) {

	for _, nxt := range g[cur] {
		// If its' neighbor has not been visited.
		if seen[nxt] == 0 {
			// Mark its neighbor to been visited.
			seen[nxt]++
			// Then deep search the neighbor.
			dfs(nxt, g, seen)
		}
	}
}
