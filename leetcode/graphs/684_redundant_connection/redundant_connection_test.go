package redundant_connection

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{
			name:  "triangle",
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			name:  "cycle of 4",
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			want:  []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRedundantConnection(tt.edges)
			if len(got) != 2 || got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
