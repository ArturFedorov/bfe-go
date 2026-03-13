package find_minimum_in_rotated_sorted_array

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"rotated 3 times", []int{3, 4, 5, 1, 2}, 1},
		{"rotated 4 times", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"not rotated", []int{11, 13, 15, 17}, 11},
		{"two elements", []int{2, 1}, 1},
		{"single element", []int{1}, 1},
		{"rotated once", []int{2, 3, 4, 5, 1}, 1},
		{"sorted ascending", []int{1, 2, 3, 4, 5}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.nums)
			if result != tt.expected {
				t.Errorf("findMin(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
