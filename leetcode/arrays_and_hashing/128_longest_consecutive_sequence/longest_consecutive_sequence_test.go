package longest_consecutive_sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
		if got != 4 {
			t.Errorf("got %d, want 4", got)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
		if got != 9 {
			t.Errorf("got %d, want 9", got)
		}
	})

	t.Run("EmptyArray", func(t *testing.T) {
		got := longestConsecutive([]int{})
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("SingleElement", func(t *testing.T) {
		got := longestConsecutive([]int{42})
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("Duplicates", func(t *testing.T) {
		got := longestConsecutive([]int{1, 2, 2, 3})
		if got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})
}
