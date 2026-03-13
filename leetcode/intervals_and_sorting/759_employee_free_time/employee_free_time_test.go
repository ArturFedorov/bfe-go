package employee_free_time

import (
	"reflect"
	"testing"
)

func TestEmployeeFreeTime(t *testing.T) {
	tests := []struct {
		name     string
		schedule [][]Interval
		want     []Interval
	}{
		{
			name: "three employees with one gap",
			schedule: [][]Interval{
				{{1, 2}, {5, 6}},
				{{1, 3}},
				{{4, 10}},
			},
			want: []Interval{{3, 4}},
		},
		{
			name: "three employees with two gaps",
			schedule: [][]Interval{
				{{1, 3}, {6, 7}},
				{{2, 4}},
				{{2, 5}, {9, 12}},
			},
			want: []Interval{{5, 6}, {7, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := employeeFreeTime(tt.schedule)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeFreeTime(%v) = %v, want %v", tt.schedule, got, tt.want)
			}
		})
	}
}
