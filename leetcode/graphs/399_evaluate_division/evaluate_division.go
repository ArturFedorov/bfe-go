package evaluate_division

type Vertex struct {
	Name  string
	Value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	vertexMap := make(map[string][]Vertex)

	for idx, equation := range equations {
		v1 := equation[0]
		v2 := equation[1]

		vertexMap[v1] = append(vertexMap[v1], Vertex{v2, values[idx]})
		vertexMap[v2] = append(vertexMap[v2], Vertex{v1, 1.0 / values[idx]})
	}

	res := make([]float64, len(queries))

	for idx, query := range queries {
		v1 := query[0]
		v2 := query[1]
		visited := make(map[string]bool)

		if _, ok := vertexMap[v1]; !ok {
			res[idx] = -1.0
		} else {
			res[idx], _ = dfs(v1, v2, vertexMap, visited)
		}
	}

	return res
}

func dfs(start, end string, vertexMap map[string][]Vertex, visited map[string]bool) (float64, bool) {
	if start == end {
		return 1.0, true
	}

	for _, next := range vertexMap[start] {
		if visited[next.Name] {
			continue
		}

		visited[next.Name] = true
		if value, ok := dfs(next.Name, end, vertexMap, visited); ok {
			return value * next.Value, true
		}
	}

	return -1.0, false
}
