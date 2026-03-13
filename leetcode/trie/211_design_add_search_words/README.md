# 211. Design Add and Search Words Data Structure

**Difficulty:** Medium

**Topics:** String, Depth-First Search, Design, Trie

## Description

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the `WordDictionary` class:
- `WordDictionary()` Initializes the object.
- `void AddWord(word)` Adds `word` to the data structure, it can be matched later.
- `bool Search(word)` Returns `true` if there is any string in the data structure that matches `word` or `false` otherwise. `word` may contain dots `'.'` where dots can be matched with any letter.

## Examples

**Example 1:**
```
Input:
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output: [null,null,null,null,false,true,true,true]
```

## Constraints

- `1 <= word.length <= 25`
- `word` in `addWord` consists of lowercase English letters.
- `word` in `search` consists of `'.'` or lowercase English letters.
- There will be at most 3 dots in `word` for `search` queries.
- At most `10^4` calls will be made to `addWord` and `search`.

## Approach Hints

1. **Trie with DFS for wildcards:** Build a standard trie for `addWord`. For `search`, when encountering `.`, branch out to all children via DFS.
2. **Optimization:** Limit branching on `.` by checking only non-nil children.
3. **Time:** O(m) for addWord, O(26^d * m) for search where d is number of dots. **Space:** O(total characters inserted).

## Related Problems

- [208. Implement Trie](../208_implement_trie/)
- [212. Word Search II](../212_word_search_ii/)

## Google Follow-ups

- How would you support deletion?
- What if the wildcard can match multiple characters (like `*` in regex)?
- How would you optimize for frequent searches with many dots?
