package find_first_and_last_position

import "testing"

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [2]int
	}{
		{"found range", []int{5, 7, 7, 8, 8, 10}, 8, [2]int{3, 4}},
		{"not found", []int{5, 7, 7, 8, 8, 10}, 6, [2]int{-1, -1}},
		{"empty array", []int{}, 0, [2]int{-1, -1}},
		{"single element found", []int{1}, 1, [2]int{0, 0}},
		{"single element not found", []int{1}, 2, [2]int{-1, -1}},
		{"all same found", []int{2, 2, 2}, 2, [2]int{0, 2}},
		{"target at start", []int{1, 1, 2, 3}, 1, [2]int{0, 1}},
		{"target at end", []int{1, 2, 3, 3}, 3, [2]int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchRange(tt.nums, tt.target)
			if len(result) != 2 || result[0] != tt.expected[0] || result[1] != tt.expected[1] {
				t.Errorf("searchRange(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
