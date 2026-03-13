package n_queens

import (
	"reflect"
	"sort"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		wantCount int
		want      [][]string
	}{
		{
			name:      "n=4 has 2 solutions",
			n:         4,
			wantCount: 2,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name:      "n=1",
			n:         1,
			wantCount: 1,
			want:      [][]string{{"Q"}},
		},
		{
			name:      "n=2 has no solutions",
			n:         2,
			wantCount: 0,
			want:      nil,
		},
		{
			name:      "n=8 has 92 solutions",
			n:         8,
			wantCount: 92,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.n)
			if len(got) != tt.wantCount {
				t.Errorf("solveNQueens(%d) returned %d solutions, want %d", tt.n, len(got), tt.wantCount)
				return
			}
			if tt.want != nil {
				sortBoards(got)
				sortBoards(tt.want)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("solveNQueens(%d) = %v, want %v", tt.n, got, tt.want)
				}
			}
		})
	}
}

func sortBoards(boards [][]string) {
	sort.Slice(boards, func(i, j int) bool {
		for k := 0; k < len(boards[i]); k++ {
			if boards[i][k] != boards[j][k] {
				return boards[i][k] < boards[j][k]
			}
		}
		return false
	})
}
