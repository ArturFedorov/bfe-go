package combination_sum

import (
	"reflect"
	"sort"
	"testing"
)

func sortResult(result [][]int) {
	for _, combo := range result {
		sort.Ints(combo)
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

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "two combinations",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "three combinations",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "no combination",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			if got == nil {
				got = [][]int{}
			}
			sortResult(got)
			sortResult(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum(%v, %d) = %v, want %v", tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}
