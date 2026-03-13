# 208. Implement Trie (Prefix Tree)

**Difficulty:** Medium

**Topics:** Hash Table, String, Design, Trie

## Description

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:
- `Trie()` Initializes the trie object.
- `void Insert(word string)` Inserts the string `word` into the trie.
- `bool Search(word string)` Returns `true` if the string `word` is in the trie (i.e., was inserted before), and `false` otherwise.
- `bool StartsWith(prefix string)` Returns `true` if there is a previously inserted string `word` that has the prefix `prefix`, and `false` otherwise.

## Examples

**Example 1:**
```
Input:
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output: [null, null, true, false, true, null, true]
```

## Constraints

- `1 <= word.length, prefix.length <= 2000`
- `word` and `prefix` consist only of lowercase English letters.
- At most `3 * 10^4` calls in total will be made to `insert`, `search`, and `startsWith`.

## Approach Hints

1. **Node structure:** Each node has up to 26 children (one per letter) and an `isEnd` flag.
2. **Insert:** Walk down the trie creating nodes as needed, mark the last node as end.
3. **Search vs StartsWith:** Search requires the final node to be marked as end; StartsWith only requires the prefix path to exist.
4. **Time:** O(m) per operation where m is key length. **Space:** O(alphabet_size * m * n) for n keys.

## Related Problems

- [211. Design Add and Search Words Data Structure](../211_design_add_search_words/)
- [212. Word Search II](../212_word_search_ii/)
- [642. Design Search Autocomplete System](../../design/642_design_search_autocomplete_system/)

## Google Follow-ups

- How would you implement delete?
- How would you implement a compressed trie (radix tree)?
- How would you count the number of words with a given prefix?
