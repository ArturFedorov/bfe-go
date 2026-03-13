package generate_parentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{"n=3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"n=1", 1, []string{"()"}},
		{"n=0", 0, []string{}},
		{"n=2", 2, []string{"(())", "()()"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateParenthesis(tt.n)
			if len(result) != len(tt.expected) {
				t.Errorf("generateParenthesis(%d) returned %d results, want %d", tt.n, len(result), len(tt.expected))
				return
			}
			expectedSet := make(map[string]bool)
			for _, s := range tt.expected {
				expectedSet[s] = true
			}
			for _, s := range result {
				if !expectedSet[s] {
					t.Errorf("generateParenthesis(%d) contains unexpected %q", tt.n, s)
				}
			}
		})
	}
}
