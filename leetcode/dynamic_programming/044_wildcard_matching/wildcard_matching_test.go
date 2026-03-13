package wildcard_matching

import "testing"

func TestIsMatchWild(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		{"no match single char", "aa", "a", false},
		{"star matches all", "aa", "*", true},
		{"question mark mismatch", "cb", "?a", false},
		{"star with prefix suffix", "adceb", "*a*b", true},
		{"empty both", "", "", true},
		{"empty string star pattern", "", "*", true},
		{"empty string non-star", "", "a", false},
		{"exact match", "abc", "abc", true},
		{"question matches any", "a", "?", true},
		{"complex pattern", "abcde", "a*e", true},
		{"star matches empty middle", "ab", "a*b", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatchWild(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatchWild(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}
