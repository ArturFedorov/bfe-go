package meeting_rooms_ii

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "two rooms needed",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      2,
		},
		{
			name:      "one room sufficient",
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      1,
		},
		{
			name:      "back to back meetings",
			intervals: [][]int{{0, 1}, {1, 2}},
			want:      1,
		},
		{
			name:      "partial overlap",
			intervals: [][]int{{1, 5}, {2, 3}, {3, 4}},
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMeetingRooms(tt.intervals)
			if got != tt.want {
				t.Errorf("minMeetingRooms(%v) = %d, want %d", tt.intervals, got, tt.want)
			}
		})
	}
}
