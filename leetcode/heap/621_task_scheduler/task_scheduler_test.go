package task_scheduler

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{
			name:  "with idle slots",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			want:  8,
		},
		{
			name:  "no cooling needed",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     0,
			want:  6,
		},
		{
			name:  "many tasks with cooling",
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     2,
			want:  16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leastInterval(tt.tasks, tt.n)
			if got != tt.want {
				t.Errorf("leastInterval(%v, %d) = %d, want %d", tt.tasks, tt.n, got, tt.want)
			}
		})
	}
}
