# 198. House Robber

**Difficulty:** Medium

**Topics:** Array, Dynamic Programming

## Description

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight **without alerting the police**.

## Examples

**Example 1:**
```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3). Total = 1 + 3 = 4.
```

**Example 2:**
```
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1). Total = 2 + 9 + 1 = 12.
```

## Constraints

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 400`

## Approach Hints

1. **DP:** `dp[i] = max(dp[i-1], dp[i-2] + nums[i])` — either skip or rob house `i`.
2. **Space-optimized:** Only need the last two values.
3. **Top-down:** Recursive with memoization.

## Related Problems

- [213. House Robber II](https://leetcode.com/problems/house-robber-ii/)
- [337. House Robber III](https://leetcode.com/problems/house-robber-iii/)
- [740. Delete and Earn](https://leetcode.com/problems/delete-and-earn/)

## Google Follow-ups

- What if the houses are arranged in a circle (House Robber II)?
- What if the houses form a binary tree (House Robber III)?
- What if you must rob at least one house?
