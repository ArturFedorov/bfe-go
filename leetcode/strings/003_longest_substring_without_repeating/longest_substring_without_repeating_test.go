package longest_substring_without_repeating

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := lengthOfLongestSubstring("abcabcbb")
		if got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})

	t.Run("AllSame", func(t *testing.T) {
		got := lengthOfLongestSubstring("bbbbb")
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("Example3", func(t *testing.T) {
		got := lengthOfLongestSubstring("pwwkew")
		if got != 3 {
			t.Errorf("got %d, want 3", got)
		}
	})

	t.Run("EmptyString", func(t *testing.T) {
		got := lengthOfLongestSubstring("")
		if got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})

	t.Run("TwoDistinct", func(t *testing.T) {
		got := lengthOfLongestSubstring("au")
		if got != 2 {
			t.Errorf("got %d, want 2", got)
		}
	})

	t.Run("SingleChar", func(t *testing.T) {
		got := lengthOfLongestSubstring("a")
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})
}
