package median_of_two_sorted_arrays

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{"odd total", []int{1, 3}, []int{2}, 2.0},
		{"even total", []int{1, 2}, []int{3, 4}, 2.5},
		{"empty first", []int{}, []int{1}, 1.0},
		{"empty second", []int{2}, []int{}, 2.0},
		{"single elements", []int{1}, []int{2}, 1.5},
		{"same elements", []int{1, 1}, []int{1, 1}, 1.0},
		{"negative numbers", []int{-3, -1}, []int{-2, 0}, -1.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(result-tt.expected) > 1e-5 {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v",
					tt.nums1, tt.nums2, result, tt.expected)
			}
		})
	}
}
