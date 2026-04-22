package design_search_autocomplete

import "sort"

type TrieNode struct {
	children  map[rune]*TrieNode
	sentences map[string]int // frequency count
}

type AutocompleteSystem struct {
	root    *TrieNode
	current *TrieNode
	input   string
}

type pair struct {
	sentence string
	freq     int
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	root := &TrieNode{
		children:  make(map[rune]*TrieNode),
		sentences: make(map[string]int),
	}

	sys := AutocompleteSystem{root: root, current: root}

	for i, s := range sentences {
		sys.addSentence(s, times[i])
	}

	return sys
}

func (a *AutocompleteSystem) addSentence(s string, freq int) {
	node := a.root
	for _, ch := range s {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children:  make(map[rune]*TrieNode),
				sentences: make(map[string]int),
			}
		}

		node = node.children[ch]
		node.sentences[s] += freq
	}
}
func (a *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		a.addSentence(a.input, 1)
		a.input = ""
		a.current = a.root
		return []string{}
	}

	a.input += string(c)

	if a.current != nil {
		a.current = a.current.children[rune(c)]
	}

	if a.current == nil {
		return []string{}
	}

	var pairs []pair
	for s, f := range a.current.sentences {
		pairs = append(pairs, pair{s, f})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].freq != pairs[j].freq {
			return pairs[i].freq > pairs[j].freq
		}

		return pairs[i].sentence < pairs[j].sentence
	})

	var result []string
	for i := 0; i < len(pairs) && i < 3; i++ {
		result = append(result, pairs[i].sentence)
	}

	return result
}
