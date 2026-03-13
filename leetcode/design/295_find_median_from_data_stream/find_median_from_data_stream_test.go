package find_median_from_data_stream

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	almostEqual := func(a, b float64) bool {
		return math.Abs(a-b) < 1e-5
	}

	t.Run("two elements median is average", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(1)
		mf.AddNum(2)
		if got := mf.FindMedian(); !almostEqual(got, 1.5) {
			t.Errorf("FindMedian() = %f, want 1.5", got)
		}
	})

	t.Run("three elements median is middle", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(1)
		mf.AddNum(2)
		mf.AddNum(3)
		if got := mf.FindMedian(); !almostEqual(got, 2.0) {
			t.Errorf("FindMedian() = %f, want 2.0", got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(42)
		if got := mf.FindMedian(); !almostEqual(got, 42.0) {
			t.Errorf("FindMedian() = %f, want 42.0", got)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(-5)
		mf.AddNum(-3)
		mf.AddNum(-1)
		if got := mf.FindMedian(); !almostEqual(got, -3.0) {
			t.Errorf("FindMedian() = %f, want -3.0", got)
		}
		mf.AddNum(-7)
		if got := mf.FindMedian(); !almostEqual(got, -4.0) {
			t.Errorf("FindMedian() = %f, want -4.0", got)
		}
	})
}
