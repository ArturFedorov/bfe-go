package subarray_sums_divisible_by_k

import (
	"testing"
)

func TestSubarraysDivByK(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)
		if got != 7 {
			t.Errorf("got %d, want 7", got)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := subarraysDivByK([]int{5}, 9)
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("AllDivisible", func(t *testing.T) {
		got := subarraysDivByK([]int{5, 10, 15}, 5)
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})

	t.Run("SingleElementDivisible", func(t *testing.T) {
		got := subarraysDivByK([]int{0}, 5)
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		got := subarraysDivByK([]int{-5, 10, -3, 3}, 5)
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})

	t.Run("AllZeros", func(t *testing.T) {
		got := subarraysDivByK([]int{0, 0, 0}, 3)
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		got := subarraysDivByK([]int{1, 2, 3}, 7)
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("KEqualsTwo", func(t *testing.T) {
		got := subarraysDivByK([]int{2, -2, 2, -4}, 2)
		if got != 10 {
			t.Errorf("got %d, want 10", got)
		}
	})

	t.Run("LargeK", func(t *testing.T) {
		got := subarraysDivByK([]int{1, 2, 3}, 10000)
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})
}
