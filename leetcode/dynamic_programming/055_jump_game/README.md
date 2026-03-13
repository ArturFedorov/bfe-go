# 55. Jump Game

**Difficulty:** Medium

**Topics:** Array, Dynamic Programming, Greedy

## Description

You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.

Return `true` if you can reach the last index, or `false` otherwise.

## Examples

**Example 1:**
```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**
```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3, whose value is 0, and you can never reach the last index.
```

## Constraints

- `1 <= nums.length <= 10^4`
- `0 <= nums[i] <= 10^5`

## Approach Hints

1. **Greedy:** Track the farthest reachable index. Iterate and update; if current index exceeds farthest, return false.
2. **DP (bottom-up):** Mark each index as good/bad starting from the end.
3. **DP (top-down):** Recursively check if the last index is reachable with memoization.

## Related Problems

- [45. Jump Game II](https://leetcode.com/problems/jump-game-ii/)
- [1306. Jump Game III](https://leetcode.com/problems/jump-game-iii/)
- [1345. Jump Game IV](https://leetcode.com/problems/jump-game-iv/)

## Google Follow-ups

- What is the minimum number of jumps to reach the end?
- What if you can also jump backwards?
- How would you handle this on a circular array?
