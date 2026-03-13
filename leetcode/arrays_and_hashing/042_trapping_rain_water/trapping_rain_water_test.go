package trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := trap([]int{4, 2, 0, 3, 2, 5})
		if got != 9 {
			t.Errorf("got %d, want 9", got)
		}
	})

	t.Run("EmptyArray", func(t *testing.T) {
		got := trap([]int{})
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("SingleElement", func(t *testing.T) {
		got := trap([]int{3})
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("Descending", func(t *testing.T) {
		got := trap([]int{5, 4, 3, 2, 1})
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("Ascending", func(t *testing.T) {
		got := trap([]int{1, 2, 3, 4, 5})
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})
}
