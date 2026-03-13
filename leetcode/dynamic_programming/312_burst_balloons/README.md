# 312. Burst Balloons

**Difficulty:** Hard

**Topics:** Array, Dynamic Programming

## Description

You are given `n` balloons, indexed from `0` to `n - 1`. Each balloon is painted with a number on it represented by an array `nums`. You are asked to burst all the balloons.

If you burst the `i`th balloon, you will get `nums[i - 1] * nums[i] * nums[i + 1]` coins. If `i - 1` or `i + 1` goes out of bounds of the array, then treat it as if there is a balloon with a `1` painted on it.

Return the maximum coins you can collect by bursting the balloons wisely.

## Examples

**Example 1:**
```
Input: nums = [3,1,5,8]
Output: 167
Explanation:
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 15 + 120 + 24 + 8 = 167
```

**Example 2:**
```
Input: nums = [1,5]
Output: 10
```

## Constraints

- `n == nums.length`
- `1 <= n <= 300`
- `0 <= nums[i] <= 100`

## Approach Hints

1. **Interval DP:** Think about which balloon to burst **last** in a range. `dp[i][j]` = max coins from bursting all balloons between `i` and `j`.
2. **Key insight:** If balloon `k` is the last burst in range `(i, j)`, then `dp[i][j] = max(dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j])`.
3. **Pad array:** Add `1` at both ends to simplify boundary handling.

## Related Problems

- [1000. Minimum Cost to Merge Stones](https://leetcode.com/problems/minimum-cost-to-merge-stones/)
- [1039. Minimum Score Triangulation of Polygon](https://leetcode.com/problems/minimum-score-triangulation-of-polygon/)

## Google Follow-ups

- Can you explain why we think about the "last balloon to burst" instead of the first?
- What is the time and space complexity of the interval DP approach?
- How would you reconstruct the optimal bursting order?
