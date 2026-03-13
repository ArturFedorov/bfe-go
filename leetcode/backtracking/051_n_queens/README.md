# 51. N-Queens

**Difficulty:** Hard

**Topics:** Array, Backtracking

## Description

The n-queens puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

Given an integer `n`, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.

## Examples

**Example 1:**
```
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
```

**Example 2:**
```
Input: n = 1
Output: [["Q"]]
```

## Constraints

- `1 <= n <= 9`

## Approach Hints

1. **Backtracking row by row:** Place queens one row at a time, checking column, diagonal, and anti-diagonal conflicts.
2. **Use sets/bitmasks:** Track occupied columns, diagonals (row-col), and anti-diagonals (row+col) for O(1) conflict checking.
3. **Time:** O(n!). **Space:** O(n^2) for board representation.

## Related Problems

- [52. N-Queens II](https://leetcode.com/problems/n-queens-ii/)
- [46. Permutations](../046_permutations/)
- [79. Word Search](../079_word_search/)

## Google Follow-ups

- Can you solve N-Queens using bit manipulation for maximum efficiency?
- How would you find just one valid solution instead of all solutions?
- What's the time complexity analysis? How does it compare to brute force?
