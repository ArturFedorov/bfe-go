package word_search

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name:  "ABCCED exists",
			board: board,
			word:  "ABCCED",
			want:  true,
		},
		{
			name:  "SEE exists",
			board: board,
			word:  "SEE",
			want:  true,
		},
		{
			name:  "ABCB does not exist",
			board: board,
			word:  "ABCB",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("exist(board, %q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
