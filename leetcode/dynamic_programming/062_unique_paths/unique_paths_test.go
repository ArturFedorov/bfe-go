package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{"3x7 grid", 3, 7, 28},
		{"3x2 grid", 3, 2, 3},
		{"1x1 grid", 1, 1, 1},
		{"7x3 grid", 7, 3, 28},
		{"1x5 grid", 1, 5, 1},
		{"5x1 grid", 5, 1, 1},
		{"2x2 grid", 2, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePaths(tt.m, tt.n)
			if result != tt.expected {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, result, tt.expected)
			}
		})
	}
}
