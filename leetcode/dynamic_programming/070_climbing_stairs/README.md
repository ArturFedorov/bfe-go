# 70. Climbing Stairs

**Difficulty:** Easy

**Topics:** Math, Dynamic Programming, Memoization

## Description

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

## Examples

**Example 1:**
```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

**Example 2:**
```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

## Constraints

- `1 <= n <= 45`

## Approach Hints

1. **DP:** `dp[i] = dp[i-1] + dp[i-2]` — this is the Fibonacci sequence.
2. **Space-optimized:** Only keep the last two values.
3. **Matrix exponentiation:** Compute the n-th Fibonacci number in O(log n) time.

## Related Problems

- [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)
- [746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)
- [1137. N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/)

## Google Follow-ups

- What if you can climb 1, 2, or 3 steps at a time?
- What if certain steps are broken and cannot be stepped on?
- Can you solve it in O(log n) time using matrix exponentiation?
