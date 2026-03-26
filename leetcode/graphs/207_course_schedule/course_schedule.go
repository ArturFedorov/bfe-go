package course_schedule

func dfs(course int, visitSet map[int]bool, adjList map[int][]int) bool {
	if ok := visitSet[course]; ok {
		return false
	}

	if len(adjList[course]) == 0 {
		return true
	}

	visitSet[course] = true

	for _, adj := range adjList[course] {
		if !dfs(adj, visitSet, adjList) {
			return false
		}
	}

	visitSet[course] = false
	adjList[course] = []int{}
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)

	for _, edge := range prerequisites {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
	}

	visitSet := make(map[int]bool)
	for i := range numCourses {
		if !dfs(i, visitSet, adjList) {
			return false
		}
	}

	return true
}
