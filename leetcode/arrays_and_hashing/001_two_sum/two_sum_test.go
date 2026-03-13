package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := twoSum([]int{2, 7, 11, 15}, 9)
		want := []int{0, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := twoSum([]int{3, 2, 4}, 6)
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example3", func(t *testing.T) {
		got := twoSum([]int{3, 3}, 6)
		want := []int{0, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		got := twoSum([]int{-1, -2, -3, -4, -5}, -8)
		want := []int{2, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("LargeTarget", func(t *testing.T) {
		got := twoSum([]int{1, 2, 3, 4, 5}, 9)
		want := []int{3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
