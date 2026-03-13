package design_add_search_words

import "testing"

func TestWordDictionary(t *testing.T) {
	t.Run("exact search", func(t *testing.T) {
		wd := Constructor()
		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")

		if !wd.Search("bad") {
			t.Error("Search(\"bad\") = false, want true")
		}

		if wd.Search("pad") {
			t.Error("Search(\"pad\") = true, want false")
		}
	})

	t.Run("wildcard dot matches any character", func(t *testing.T) {
		wd := Constructor()
		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")

		if !wd.Search(".ad") {
			t.Error("Search(\".ad\") = false, want true")
		}

		if !wd.Search("b.d") {
			t.Error("Search(\"b.d\") = false, want true")
		}

		if !wd.Search("b..") {
			t.Error("Search(\"b..\") = false, want true")
		}
	})

	t.Run("no match with wildcard", func(t *testing.T) {
		wd := Constructor()
		wd.AddWord("bad")

		if wd.Search("b.") {
			t.Error("Search(\"b.\") = true, want false (wrong length)")
		}

		if wd.Search("....") {
			t.Error("Search(\"....\") = true, want false (wrong length)")
		}
	})

	t.Run("multiple words different lengths", func(t *testing.T) {
		wd := Constructor()
		wd.AddWord("a")
		wd.AddWord("ab")

		if !wd.Search("a") {
			t.Error("Search(\"a\") = false, want true")
		}

		if !wd.Search(".") {
			t.Error("Search(\".\") = false, want true")
		}

		if !wd.Search("ab") {
			t.Error("Search(\"ab\") = false, want true")
		}

		if !wd.Search(".b") {
			t.Error("Search(\".b\") = false, want true")
		}
	})
}
