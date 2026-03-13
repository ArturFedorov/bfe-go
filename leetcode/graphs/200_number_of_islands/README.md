# 200. Number of Islands

**Difficulty:** Medium

**Topics:** Array, Depth-First Search, Breadth-First Search, Union Find, Matrix

## Description

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Examples

### Example 1
```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

### Example 2
```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

## Constraints

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## Approach Hints

1. **DFS/BFS flood fill:** Iterate through the grid. When you find a '1', increment count and flood fill (mark all connected '1's as visited).
2. **Union Find:** Union adjacent land cells. Count the number of distinct sets at the end.
3. **In-place marking:** Modify the grid to mark visited cells (change '1' to '0') to avoid extra space.

## Related Problems

- 130. Surrounded Regions
- 286. Walls and Gates
- 305. Number of Islands II
- 695. Max Area of Island

## Google Follow-ups

- How would you count islands in a stream of updates (cells turning to land)?
- What if the grid is too large to fit in memory?
- How would you count islands if diagonal connections also count?
