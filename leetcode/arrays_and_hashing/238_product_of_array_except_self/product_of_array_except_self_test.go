package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := productExceptSelf([]int{1, 2, 3, 4})
		want := []int{24, 12, 8, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := productExceptSelf([]int{-1, 1, 0, -3, 3})
		want := []int{0, 0, 9, 0, 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("TwoElements", func(t *testing.T) {
		got := productExceptSelf([]int{2, 3})
		want := []int{3, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("WithZero", func(t *testing.T) {
		got := productExceptSelf([]int{0, 1, 2, 3})
		want := []int{6, 0, 0, 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("AllOnes", func(t *testing.T) {
		got := productExceptSelf([]int{1, 1, 1, 1})
		want := []int{1, 1, 1, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		got := productExceptSelf([]int{-1, -2, -3})
		want := []int{6, 3, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
