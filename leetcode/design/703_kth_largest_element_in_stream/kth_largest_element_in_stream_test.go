package kth_largest_element_in_stream

import "testing"

func TestKthLargest(t *testing.T) {
	t.Run("k=3 [4,5,8,2] add sequence", func(t *testing.T) {
		kl := Constructor(3, []int{4, 5, 8, 2})

		tests := []struct {
			val  int
			want int
		}{
			{3, 4},
			{5, 5},
			{10, 5},
			{9, 8},
			{4, 8},
		}

		for _, tt := range tests {
			if got := kl.Add(tt.val); got != tt.want {
				t.Errorf("Add(%d) = %d, want %d", tt.val, got, tt.want)
			}
		}
	})

	t.Run("k=1 single element", func(t *testing.T) {
		kl := Constructor(1, []int{})
		if got := kl.Add(1); got != 1 {
			t.Errorf("Add(1) = %d, want 1", got)
		}
		if got := kl.Add(2); got != 2 {
			t.Errorf("Add(2) = %d, want 2", got)
		}
	})

	t.Run("k=2 with negative numbers", func(t *testing.T) {
		kl := Constructor(2, []int{-1, -2, -3})
		if got := kl.Add(-4); got != -2 {
			t.Errorf("Add(-4) = %d, want -2", got)
		}
		if got := kl.Add(0); got != -1 {
			t.Errorf("Add(0) = %d, want -1", got)
		}
	})
}
