package word_search_ii

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			name: "find eat and oath",
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"eat", "oath"},
		},
		{
			name: "no words found",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words: []string{"abcb"},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWords(tt.board, tt.words)
			if got == nil {
				got = []string{}
			}
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords(board, %v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}
