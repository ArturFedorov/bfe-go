package letter_combinations

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := letterCombinations("23")
		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		sort.Strings(got)
		sort.Strings(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("EmptyString", func(t *testing.T) {
		got := letterCombinations("")
		if len(got) != 0 {
			t.Errorf("got %v, want []", got)
		}
	})

	t.Run("SingleDigit", func(t *testing.T) {
		got := letterCombinations("2")
		want := []string{"a", "b", "c"}
		sort.Strings(got)
		sort.Strings(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("DigitWith4Letters", func(t *testing.T) {
		got := letterCombinations("7")
		want := []string{"p", "q", "r", "s"}
		sort.Strings(got)
		sort.Strings(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("ThreeDigits", func(t *testing.T) {
		got := letterCombinations("234")
		if len(got) != 27 {
			t.Errorf("got %d combinations, want 27", len(got))
		}
	})
}
