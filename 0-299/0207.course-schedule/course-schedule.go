package problem0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	// use a map go represent the graph
	graph := make(map[int][]int)

	// build the graph
	for _, pair := range prerequisites {
		course := pair[0]
		preCourse := pair[1]
		graph[course] = append(graph[course], preCourse)
	}

	// a list to denote node traversing status:
	// 0: default, not visited
	// 1: visiting
	// 2: visited
	visited := make([]int, numCourses)

	// call topological dfs on this every course
	for i := 0; i < numCourses; i++ {
		if dfsTopological(i, graph, visited) {
			return false
		}
	}

	return true
}

// topological traversing with DFS
func dfsTopological(course int, graph map[int][]int, visited []int) bool {
	// 1: visiting
	if visited[course] == 1 {
		return true
	}
	// 2: visited
	if visited[course] == 2 {
		return false
	}
	// 0: not visited
	visited[course] = 1

	preCourses, ok := graph[course]
	if ok {
		for _, pre := range preCourses {
			if dfsTopological(pre, graph, visited) {
				return true
			}
		}
	}
	visited[course] = 2
	return false
}
