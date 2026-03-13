package design_search_autocomplete

import "testing"

func equalStr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAutocompleteSystem(t *testing.T) {
	t.Run("basic input sequence", func(t *testing.T) {
		sentences := []string{"i love you", "island", "iroman", "i love leetcode"}
		times := []int{5, 3, 2, 2}
		ac := Constructor(sentences, times)

		got := ac.Input('i')
		want := []string{"i love you", "island", "i love leetcode"}
		if !equalStr(got, want) {
			t.Errorf("Input('i') = %v, want %v", got, want)
		}

		got = ac.Input(' ')
		want = []string{"i love you", "i love leetcode"}
		if !equalStr(got, want) {
			t.Errorf("Input(' ') = %v, want %v", got, want)
		}
	})

	t.Run("# ends sentence", func(t *testing.T) {
		sentences := []string{"abc", "abd"}
		times := []int{3, 2}
		ac := Constructor(sentences, times)

		ac.Input('a')
		got := ac.Input('#')
		if len(got) != 0 {
			t.Errorf("Input('#') = %v, want empty", got)
		}
	})

	t.Run("new sentence appears in results", func(t *testing.T) {
		sentences := []string{"abc"}
		times := []int{1}
		ac := Constructor(sentences, times)

		ac.Input('x')
		ac.Input('y')
		ac.Input('#')

		got := ac.Input('x')
		want := []string{"xy"}
		if !equalStr(got, want) {
			t.Errorf("Input('x') after adding 'xy' = %v, want %v", got, want)
		}
		ac.Input('#')
	})

	t.Run("top 3 by frequency", func(t *testing.T) {
		sentences := []string{"a", "b", "c", "d"}
		times := []int{4, 3, 2, 1}
		ac := Constructor(sentences, times)

		got := ac.Input('#')
		if len(got) != 0 {
			t.Errorf("Input('#') = %v, want empty", got)
		}
	})
}
