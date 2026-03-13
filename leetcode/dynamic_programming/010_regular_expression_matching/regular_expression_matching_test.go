package regular_expression_matching

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		{"no match single char", "aa", "a", false},
		{"star matches repeated char", "aa", "a*", true},
		{"dot star matches any", "ab", ".*", true},
		{"star can match zero", "aab", "c*a*b", true},
		{"exact match", "abc", "abc", true},
		{"empty both", "", "", true},
		{"empty string with star pattern", "", "a*", true},
		{"empty string with complex star", "", "a*b*c*", true},
		{"dot matches single", "a", ".", true},
		{"no match different char", "a", "b", false},
		{"star matches zero occurrences", "b", "a*b", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatch(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}
