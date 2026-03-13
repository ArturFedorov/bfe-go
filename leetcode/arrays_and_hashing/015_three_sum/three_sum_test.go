package three_sum

import (
	"reflect"
	"sort"
	"testing"
)

func sortResult(result [][]int) {
	for _, triplet := range result {
		sort.Ints(triplet)
	}
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i]) && k < len(result[j]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return len(result[i]) < len(result[j])
	})
}

func TestThreeSum(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		got := threeSum(nums)
		want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
		sortResult(got)
		sortResult(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("NoTriplets", func(t *testing.T) {
		nums := []int{0, 1, 1}
		got := threeSum(nums)
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("AllZeros", func(t *testing.T) {
		nums := []int{0, 0, 0}
		got := threeSum(nums)
		want := [][]int{{0, 0, 0}}
		sortResult(got)
		sortResult(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("EmptyArray", func(t *testing.T) {
		nums := []int{}
		got := threeSum(nums)
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("TwoElements", func(t *testing.T) {
		nums := []int{1, -1}
		got := threeSum(nums)
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})
}
