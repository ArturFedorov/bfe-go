package longest_palindrome_build

import "testing"

func TestLongestPalindromeBuild(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := longestPalindromeBuild("abccccdd")
		if got != 7 {
			t.Errorf("got %d, want 7", got)
		}
	})

	t.Run("SingleChar", func(t *testing.T) {
		got := longestPalindromeBuild("a")
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("CaseSensitive", func(t *testing.T) {
		got := longestPalindromeBuild("Aa")
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})

	t.Run("AllSame", func(t *testing.T) {
		got := longestPalindromeBuild("aaaa")
		if got != 4 {
			t.Errorf("got %d, want 4", got)
		}
	})

	t.Run("AllDistinct", func(t *testing.T) {
		got := longestPalindromeBuild("abcdef")
		if got != 1 {
			t.Errorf("got %d, want 1", got)
		}
	})
}
