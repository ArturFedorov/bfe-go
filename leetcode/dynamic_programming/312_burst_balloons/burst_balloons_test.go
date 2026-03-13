package burst_balloons

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1", []int{3, 1, 5, 8}, 167},
		{"two balloons", []int{1, 5}, 10},
		{"single balloon", []int{1}, 1},
		{"single large", []int{9}, 9},
		{"three balloons", []int{1, 2, 3}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxCoins(tt.nums)
			if result != tt.expected {
				t.Errorf("maxCoins(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
