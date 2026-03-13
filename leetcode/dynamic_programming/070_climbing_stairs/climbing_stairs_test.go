package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"n=2", 2, 2},
		{"n=3", 3, 3},
		{"n=1", 1, 1},
		{"n=4", 4, 5},
		{"n=45", 45, 1836311903},
		{"n=5", 5, 8},
		{"n=10", 10, 89},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := climbStairs(tt.n)
			if result != tt.expected {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
