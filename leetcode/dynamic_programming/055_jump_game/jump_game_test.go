package jump_game

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"can reach end", []int{2, 3, 1, 1, 4}, true},
		{"stuck at zero", []int{3, 2, 1, 0, 4}, false},
		{"single element", []int{0}, true},
		{"jump over zeros", []int{2, 0, 0}, true},
		{"all zeros except first", []int{1, 0, 0}, false},
		{"large first jump", []int{5, 0, 0, 0, 0}, true},
		{"single nonzero", []int{1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			if result != tt.expected {
				t.Errorf("canJump(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
