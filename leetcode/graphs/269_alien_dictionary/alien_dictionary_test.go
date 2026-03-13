package alien_dictionary

import "testing"

func isValidAlienOrder(order string, words []string) bool {
	if order == "" {
		return false
	}
	pos := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		pos[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		found := false
		for j := 0; j < len(w1) && j < len(w2); j++ {
			if w1[j] != w2[j] {
				if pos[w1[j]] > pos[w2[j]] {
					return false
				}
				found = true
				break
			}
		}
		if !found && len(w1) > len(w2) {
			return false
		}
	}
	return true
}

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		valid bool
	}{
		{
			name:  "wrt,wrf,er,ett,rftt",
			words: []string{"wrt", "wrf", "er", "ett", "rftt"},
			valid: true,
		},
		{
			name:  "z,x",
			words: []string{"z", "x"},
			valid: true,
		},
		{
			name:  "z,x,z invalid",
			words: []string{"z", "x", "z"},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alienOrder(tt.words)
			if tt.valid {
				if got == "" || !isValidAlienOrder(got, tt.words) {
					t.Errorf("alienOrder() = %q, not a valid alien ordering for %v", got, tt.words)
				}
			} else {
				if got != "" {
					t.Errorf("alienOrder() = %q, want empty string for invalid input", got)
				}
			}
		})
	}
}
