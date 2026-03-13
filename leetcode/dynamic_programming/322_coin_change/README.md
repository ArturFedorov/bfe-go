# 322. Coin Change

**Difficulty:** Medium

**Topics:** Array, Dynamic Programming, Breadth-First Search

## Description

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

## Examples

**Example 1:**
```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
```

**Example 2:**
```
Input: coins = [2], amount = 3
Output: -1
```

**Example 3:**
```
Input: coins = [1], amount = 0
Output: 0
```

## Constraints

- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`

## Approach Hints

1. **DP (bottom-up):** `dp[i]` = min coins to make amount `i`. For each coin, `dp[i] = min(dp[i], dp[i-coin] + 1)`.
2. **BFS:** Treat amounts as nodes, each coin as an edge. BFS finds shortest path (fewest coins).
3. **Top-down:** Recursive with memoization on remaining amount.

## Related Problems

- [518. Coin Change II](https://leetcode.com/problems/coin-change-ii/)
- [983. Minimum Cost For Tickets](https://leetcode.com/problems/minimum-cost-for-tickets/)
- [1049. Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/)

## Google Follow-ups

- How many different ways can you make the amount (Coin Change II)?
- What if each coin can only be used once?
- How would you handle very large amounts efficiently?
