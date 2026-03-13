package subarray_sum_equals_k

import (
	"testing"
)

func TestSubarraySum(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := subarraySum([]int{1, 1, 1}, 2)
		if got != 2 {
			t.Errorf("got %d, want 2", got)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := subarraySum([]int{1, 2, 3}, 3)
		if got != 2 {
			t.Errorf("got %d, want 2", got)
		}
	})

	t.Run("SingleElement_Match", func(t *testing.T) {
		got := subarraySum([]int{5}, 5)
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("SingleElement_NoMatch", func(t *testing.T) {
		got := subarraySum([]int{5}, 3)
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		got := subarraySum([]int{-1, -1, 1}, 0)
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("AllZeros", func(t *testing.T) {
		got := subarraySum([]int{0, 0, 0}, 0)
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})

	t.Run("EntireArray", func(t *testing.T) {
		got := subarraySum([]int{1, 2, 3}, 6)
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		got := subarraySum([]int{1, 2, 3}, 100)
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("NegativeK", func(t *testing.T) {
		got := subarraySum([]int{-1, -1, 1}, -2)
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("MultipleOverlapping", func(t *testing.T) {
		got := subarraySum([]int{1, -1, 1, -1, 1}, 0)
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})
}
