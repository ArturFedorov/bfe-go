package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			name:      "hit -> cog with path",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			name:      "hit -> cog without cog in list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
