package permutations

import (
	"reflect"
	"sort"
	"testing"
)

func sortPerms(result [][]int) {
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i]) && k < len(result[j]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return len(result[i]) < len(result[j])
	})
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "two elements",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			sortPerms(got)
			sortPerms(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
