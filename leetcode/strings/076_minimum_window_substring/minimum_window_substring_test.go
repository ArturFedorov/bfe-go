package minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := minWindow("ADOBECODEBANC", "ABC")
		if got != "BANC" {
			t.Errorf("got %q, want %q", got, "BANC")
		}
	})

	t.Run("ExactMatch", func(t *testing.T) {
		got := minWindow("a", "a")
		if got != "a" {
			t.Errorf("got %q, want %q", got, "a")
		}
	})

	t.Run("NotPossible", func(t *testing.T) {
		got := minWindow("a", "aa")
		if got != "" {
			t.Errorf("got %q, want %q", got, "")
		}
	})

	t.Run("EmptyT", func(t *testing.T) {
		got := minWindow("abc", "")
		if got != "" {
			t.Errorf("got %q, want %q", got, "")
		}
	})

	t.Run("SLongerNoMatch", func(t *testing.T) {
		got := minWindow("abcdef", "xyz")
		if got != "" {
			t.Errorf("got %q, want %q", got, "")
		}
	})

	t.Run("DuplicateCharsInT", func(t *testing.T) {
		got := minWindow("aab", "aab")
		if got != "aab" {
			t.Errorf("got %q, want %q", got, "aab")
		}
	})
}
