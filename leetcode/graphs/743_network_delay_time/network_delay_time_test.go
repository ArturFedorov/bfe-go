package network_delay_time

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][]int
		n     int
		k     int
		want  int
	}{
		{
			name:  "4 nodes, start at 2",
			times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:     4,
			k:     2,
			want:  2,
		},
		{
			name:  "unreachable node",
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     2,
			want:  -1,
		},
		{
			name:  "2 nodes, start at 1",
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     1,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := networkDelayTime(tt.times, tt.n, tt.k)
			if got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
