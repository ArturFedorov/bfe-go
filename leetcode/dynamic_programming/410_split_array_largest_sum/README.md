# 410. Split Array Largest Sum

**Difficulty:** Hard

**Topics:** Array, Binary Search, Dynamic Programming, Greedy, Prefix Sum

## Description

Given an integer array `nums` and an integer `k`, split `nums` into `k` non-empty subarrays such that the largest sum of any subarray is minimized.

Return the minimized largest sum of the split.

A **subarray** is a contiguous part of the array.

## Examples

**Example 1:**
```
Input: nums = [7,2,5,10,8], k = 2
Output: 18
Explanation: There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8], where the largest sum is 18.
```

**Example 2:**
```
Input: nums = [1,2,3,4,5], k = 2
Output: 9
Explanation: Split into [1,2,3] and [4,5] with largest sum 9.
```

## Constraints

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 10^6`
- `1 <= k <= min(50, nums.length)`

## Approach Hints

1. **Binary Search:** Binary search on the answer (largest sum). For each candidate, greedily check if you can split into <= k subarrays.
2. **DP:** `dp[i][j]` = min largest sum splitting first `i` elements into `j` subarrays. Use prefix sums.
3. **Key insight:** The answer lies between `max(nums)` and `sum(nums)`.

## Related Problems

- [1011. Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)
- [875. Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)
- [1231. Divide Chocolate](https://leetcode.com/problems/divide-chocolate/)

## Google Follow-ups

- Can you prove why binary search works here (monotonicity)?
- What if the array can be split into non-contiguous groups?
- How would you extend this to minimize the variance of subarray sums?
