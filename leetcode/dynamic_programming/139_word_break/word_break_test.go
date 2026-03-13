package word_break

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", "leetcode", []string{"leet", "code"}, true},
		{"applepenapple", "applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"single word match", "a", []string{"a"}, true},
		{"single word no match", "b", []string{"a"}, false},
		{"empty string", "", []string{"a"}, true},
		{"repeated word", "aaaa", []string{"a", "aa"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, result, tt.expected)
			}
		})
	}
}
