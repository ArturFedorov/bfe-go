# 79. Word Search

**Difficulty:** Medium

**Topics:** Array, String, Backtracking, Matrix

## Description

Given an `m x n` grid of characters `board` and a string `word`, return `true` if `word` exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

## Examples

**Example 1:**
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

**Example 2:**
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
```

**Example 3:**
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
```

## Constraints

- `m == board.length`
- `n = board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` and `word` consists of only lowercase and uppercase English letters.

## Approach Hints

1. **DFS/Backtracking:** For each cell matching the first letter, run DFS exploring all 4 directions. Mark cells as visited during exploration and unmark on backtrack.
2. **Pruning:** Check character frequency—if the word needs more of a character than the board has, return false early.
3. **Time:** O(m * n * 3^L) where L is word length. **Space:** O(L) for recursion depth.

## Related Problems

- [212. Word Search II](../../trie/212_word_search_ii/)
- [51. N-Queens](../051_n_queens/)
- [489. Robot Room Cleaner](../489_robot_room_cleaner/)

## Google Follow-ups

- How would you search for multiple words simultaneously? (See problem 212)
- What if the board is very large? How would you optimize?
- Can you solve this iteratively instead of recursively?
