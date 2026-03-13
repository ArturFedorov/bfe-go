package find_all_anagrams

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := findAnagrams("cbaebabacd", "abc")
		want := []int{0, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := findAnagrams("abab", "ab")
		want := []int{0, 1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("EmptyString", func(t *testing.T) {
		got := findAnagrams("", "abc")
		if got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		got := findAnagrams("abcdef", "xyz")
		if got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})

	t.Run("PLongerThanS", func(t *testing.T) {
		got := findAnagrams("ab", "abc")
		if got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
}
