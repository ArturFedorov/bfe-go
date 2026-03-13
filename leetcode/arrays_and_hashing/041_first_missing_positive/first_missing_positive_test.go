package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := firstMissingPositive([]int{1, 2, 0})
		if got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := firstMissingPositive([]int{3, 4, -1, 1})
		if got != 2 {
			t.Errorf("got %d, want 2", got)
		}
	})

	t.Run("Example3", func(t *testing.T) {
		got := firstMissingPositive([]int{7, 8, 9, 11, 12})
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("SingleElement", func(t *testing.T) {
		got := firstMissingPositive([]int{1})
		if got != 2 {
			t.Errorf("got %d, want 2", got)
		}
	})

	t.Run("Consecutive", func(t *testing.T) {
		got := firstMissingPositive([]int{1, 2, 3, 4, 5})
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})

	t.Run("AllNegative", func(t *testing.T) {
		got := firstMissingPositive([]int{-1, -2, -3})
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})
}
