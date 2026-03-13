package minimum_cost_to_hire_k_workers

import (
	"math"
	"testing"
)

func TestMincostToHireWorkers(t *testing.T) {
	tests := []struct {
		name    string
		quality []int
		wage    []int
		k       int
		want    float64
	}{
		{
			name:    "two workers",
			quality: []int{10, 20, 5},
			wage:    []int{70, 50, 30},
			k:       2,
			want:    105.00000,
		},
		{
			name:    "three workers",
			quality: []int{3, 1, 10, 10, 1},
			wage:    []int{4, 8, 2, 2, 7},
			k:       3,
			want:    30.66667,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mincostToHireWorkers(tt.quality, tt.wage, tt.k)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("mincostToHireWorkers(%v, %v, %d) = %f, want %f", tt.quality, tt.wage, tt.k, got, tt.want)
			}
		})
	}
}
