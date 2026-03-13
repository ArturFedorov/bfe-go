package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("SimplePair", func(t *testing.T) {
		if !isValid("()") {
			t.Error("expected true for \"()\"")
		}
	})

	t.Run("MultiplePairs", func(t *testing.T) {
		if !isValid("()[]{}") {
			t.Error("expected true for \"()[]{}\"")
		}
	})

	t.Run("Mismatch", func(t *testing.T) {
		if isValid("(]") {
			t.Error("expected false for \"(]\"")
		}
	})

	t.Run("Interleaved", func(t *testing.T) {
		if isValid("([)]") {
			t.Error("expected false for \"([)]\"")
		}
	})

	t.Run("Nested", func(t *testing.T) {
		if !isValid("{[]}") {
			t.Error("expected true for \"{[]}\"")
		}
	})

	t.Run("EmptyString", func(t *testing.T) {
		if !isValid("") {
			t.Error("expected true for \"\"")
		}
	})

	t.Run("OnlyOpening", func(t *testing.T) {
		if isValid("(((") {
			t.Error("expected false for \"(((\"")
		}
	})

	t.Run("OnlyClosing", func(t *testing.T) {
		if isValid(")))") {
			t.Error("expected false for \")))\"")
		}
	})
}
