package max_profit_job_scheduling

import "testing"

func TestJobScheduling(t *testing.T) {
	tests := []struct {
		name      string
		startTime []int
		endTime   []int
		profit    []int
		expected  int
	}{
		{
			"example 1",
			[]int{1, 2, 3, 3},
			[]int{3, 4, 5, 6},
			[]int{50, 10, 40, 70},
			120,
		},
		{
			"example 2",
			[]int{1, 2, 3, 4, 6},
			[]int{3, 5, 10, 6, 9},
			[]int{20, 20, 100, 70, 60},
			150,
		},
		{
			"single job",
			[]int{1},
			[]int{2},
			[]int{50},
			50,
		},
		{
			"all overlapping",
			[]int{1, 1, 1},
			[]int{3, 3, 3},
			[]int{10, 20, 30},
			30,
		},
		{
			"no overlap",
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{10, 20, 30},
			60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jobScheduling(tt.startTime, tt.endTime, tt.profit)
			if result != tt.expected {
				t.Errorf("jobScheduling(%v, %v, %v) = %d, want %d",
					tt.startTime, tt.endTime, tt.profit, result, tt.expected)
			}
		})
	}
}
