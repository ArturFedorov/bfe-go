package fruit_into_baskets

import "testing"

func TestTotalFruit(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := totalFruit([]int{1, 2, 1})
		if got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := totalFruit([]int{0, 1, 2, 2})
		if got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})

	t.Run("Example3", func(t *testing.T) {
		got := totalFruit([]int{1, 2, 3, 2, 2})
		if got != 4 {
			t.Errorf("got %d, want 4", got)
		}
	})

	t.Run("LongerInput", func(t *testing.T) {
		got := totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4})
		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})

	t.Run("SingleType", func(t *testing.T) {
		got := totalFruit([]int{5, 5, 5, 5})
		if got != 4 {
			t.Errorf("got %d, want 4", got)
		}
	})

	t.Run("SingleElement", func(t *testing.T) {
		got := totalFruit([]int{7})
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})
}
