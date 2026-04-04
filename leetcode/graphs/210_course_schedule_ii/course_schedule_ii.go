package course_schedule_ii

func topologicalSortBFS(adj map[int][]int, inDegree []int) []int {
	res := make([]int, 0)
	queue := make([]int, 0)

	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
			res = append(res, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
				res = append(res, v)
			}
		}
	}

	return res
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make(map[int][]int)
	inDegree := make([]int, numCourses)

	for _, vec := range prerequisites {
		u := vec[1]
		v := vec[0]
		adj[u] = append(adj[u], v)
		inDegree[v]++
	}

	res := topologicalSortBFS(adj, inDegree)

	if len(res) != numCourses {
		return nil
	}

	return res
}
