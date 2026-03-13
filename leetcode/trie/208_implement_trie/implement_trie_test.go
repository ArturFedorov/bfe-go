package implement_trie

import "testing"

func TestTrie(t *testing.T) {
	t.Run("insert and search", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")

		if !trie.Search("apple") {
			t.Error("Search(\"apple\") = false, want true")
		}

		if trie.Search("app") {
			t.Error("Search(\"app\") = true, want false")
		}
	})

	t.Run("startsWith", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")

		if !trie.StartsWith("app") {
			t.Error("StartsWith(\"app\") = false, want true")
		}

		if trie.StartsWith("b") {
			t.Error("StartsWith(\"b\") = true, want false")
		}
	})

	t.Run("search missing word", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("hello")

		if trie.Search("world") {
			t.Error("Search(\"world\") = true, want false")
		}
	})

	t.Run("insert multiple words", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		trie.Insert("app")

		if !trie.Search("apple") {
			t.Error("Search(\"apple\") = false, want true")
		}

		if !trie.Search("app") {
			t.Error("Search(\"app\") = false, want true after inserting \"app\"")
		}

		if trie.Search("ap") {
			t.Error("Search(\"ap\") = true, want false")
		}

		if !trie.StartsWith("ap") {
			t.Error("StartsWith(\"ap\") = false, want true")
		}
	})
}
