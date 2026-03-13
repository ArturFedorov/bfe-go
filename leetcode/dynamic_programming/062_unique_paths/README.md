# 62. Unique Paths

**Difficulty:** Medium

**Topics:** Math, Dynamic Programming, Combinatorics

## Description

There is a robot on an `m x n` grid. The robot is initially located at the **top-left corner** (i.e., `grid[0][0]`). The robot tries to move to the **bottom-right corner** (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

## Examples

**Example 1:**
```
Input: m = 3, n = 7
Output: 28
```

**Example 2:**
```
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
```

## Constraints

- `1 <= m, n <= 100`

## Approach Hints

1. **DP:** `dp[i][j] = dp[i-1][j] + dp[i][j-1]` with base case `dp[0][j] = dp[i][0] = 1`.
2. **Space-optimized DP:** Use a single 1D array of size `n`.
3. **Math/Combinatorics:** Answer is `C(m+n-2, m-1)` — choose which steps go down.

## Related Problems

- [63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii/)
- [64. Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/)
- [980. Unique Paths III](https://leetcode.com/problems/unique-paths-iii/)

## Google Follow-ups

- What if some cells are blocked (obstacles)?
- Can you solve it in O(min(m,n)) space?
- What if the robot can also move diagonally?
