package longest_palindromic_substring

import "testing"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func assertValidPalindrome(t *testing.T, input, got string, maxLen int) {
	t.Helper()
	if !isPalindrome(got) {
		t.Errorf("result %q is not a palindrome", got)
	}
	if len(got) != maxLen {
		t.Errorf("got length %d, want %d", len(got), maxLen)
	}
}

func TestLongestPalindrome(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := longestPalindrome("babad")
		assertValidPalindrome(t, "babad", got, 3)
	})

	t.Run("Example2", func(t *testing.T) {
		got := longestPalindrome("cbbd")
		if got != "bb" {
			t.Errorf("got %q, want %q", got, "bb")
		}
	})

	t.Run("SingleChar", func(t *testing.T) {
		got := longestPalindrome("a")
		if got != "a" {
			t.Errorf("got %q, want %q", got, "a")
		}
	})

	t.Run("TwoDistinct", func(t *testing.T) {
		got := longestPalindrome("ac")
		assertValidPalindrome(t, "ac", got, 1)
	})

	t.Run("EntireString", func(t *testing.T) {
		got := longestPalindrome("racecar")
		if got != "racecar" {
			t.Errorf("got %q, want %q", got, "racecar")
		}
	})
}
