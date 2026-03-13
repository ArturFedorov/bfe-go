package split_array_largest_sum

import "testing"

func TestSplitArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"example 1", []int{7, 2, 5, 10, 8}, 2, 18},
		{"example 2", []int{1, 2, 3, 4, 5}, 2, 9},
		{"each element own group", []int{1, 4, 4}, 3, 4},
		{"single element", []int{10}, 1, 10},
		{"k equals length", []int{1, 2, 3}, 3, 3},
		{"all same values", []int{5, 5, 5, 5}, 2, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitArray(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("splitArray(%v, %d) = %d, want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
