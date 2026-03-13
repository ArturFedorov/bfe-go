package house_robber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1", []int{1, 2, 3, 1}, 4},
		{"example 2", []int{2, 7, 9, 3, 1}, 12},
		{"alternating best", []int{2, 1, 1, 2}, 4},
		{"single zero", []int{0}, 0},
		{"single house", []int{5}, 5},
		{"two houses", []int{1, 2}, 2},
		{"all same", []int{3, 3, 3, 3}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rob(tt.nums)
			if result != tt.expected {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
