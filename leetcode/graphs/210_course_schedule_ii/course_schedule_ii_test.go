package course_schedule_ii

import "testing"

func isValidOrder(order []int, numCourses int, prerequisites [][]int) bool {
	if len(order) != numCourses {
		return false
	}
	pos := make(map[int]int)
	for i, v := range order {
		pos[v] = i
	}
	if len(pos) != numCourses {
		return false
	}
	for _, p := range prerequisites {
		if pos[p[1]] > pos[p[0]] {
			return false
		}
	}
	return true
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
	}{
		{
			name:          "2 courses",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
		},
		{
			name:          "4 courses",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
		},
		{
			name:          "1 course no prereqs",
			numCourses:    1,
			prerequisites: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.numCourses, tt.prerequisites)
			if !isValidOrder(got, tt.numCourses, tt.prerequisites) {
				t.Errorf("findOrder() = %v, not a valid topological order", got)
			}
		})
	}
}
