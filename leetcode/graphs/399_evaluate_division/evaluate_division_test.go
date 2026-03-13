package evaluate_division

import (
	"math"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			name:      "basic equations",
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}},
			want:      []float64{6.0, 0.5, -1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEquation(tt.equations, tt.values, tt.queries)
			if len(got) != len(tt.want) {
				t.Errorf("calcEquation() returned %d results, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if tt.want[i] == -1.0 {
					if got[i] != -1.0 {
						t.Errorf("query %d: got %v, want -1.0", i, got[i])
					}
				} else if math.Abs(got[i]-tt.want[i]) > 1e-5 {
					t.Errorf("query %d: got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
