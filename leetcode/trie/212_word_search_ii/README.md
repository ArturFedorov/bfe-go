# 212. Word Search II

**Difficulty:** Hard

**Topics:** Array, String, Backtracking, Trie, Matrix

## Description

Given an `m x n` `board` of characters and a list of strings `words`, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

## Examples

**Example 1:**
```
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
```

**Example 2:**
```
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
```

## Constraints

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` is a lowercase English letter.
- `1 <= words.length <= 3 * 10^4`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- All the strings of `words` are unique.

## Approach Hints

1. **Trie + Backtracking:** Build a trie from the word list. DFS on the board while traversing the trie simultaneously. This avoids redundant searches.
2. **Pruning:** Remove words from the trie once found to avoid duplicate results and speed up future searches.
3. **Time:** O(m * n * 4^L) where L is max word length, but trie pruning makes it much faster in practice. **Space:** O(total characters in words) for trie.

## Related Problems

- [79. Word Search](../../backtracking/079_word_search/)
- [208. Implement Trie](../208_implement_trie/)
- [211. Design Add and Search Words Data Structure](../211_design_add_search_words/)

## Google Follow-ups

- How would you optimize if the board is very large but the word list is small?
- What if words can wrap around the board edges?
- How would you handle this if the board changes dynamically?
