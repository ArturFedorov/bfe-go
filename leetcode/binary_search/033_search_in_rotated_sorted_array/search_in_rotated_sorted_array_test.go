package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"found in rotated", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"not found in rotated", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single element miss", []int{1}, 0, -1},
		{"single element hit", []int{1}, 1, 0},
		{"not rotated", []int{1, 2, 3, 4, 5}, 3, 2},
		{"target at start", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"target at end", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"two elements", []int{3, 1}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
