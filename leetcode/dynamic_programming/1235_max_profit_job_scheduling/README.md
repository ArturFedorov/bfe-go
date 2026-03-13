# 1235. Maximum Profit in Job Scheduling

**Difficulty:** Hard

**Topics:** Array, Binary Search, Dynamic Programming, Sorting

## Description

We have `n` jobs, where every job is scheduled to be done from `startTime[i]` to `endTime[i]`, obtaining a profit of `profit[i]`.

You're given the `startTime`, `endTime` and `profit` arrays, return the maximum profit you can take such that there are no two jobs in the subset with overlapping time range.

If you choose a job that ends at time `X` you will be able to start another job that starts at time `X`.

## Examples

**Example 1:**
```
Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
Output: 120
Explanation: The subset chosen is the first and fourth job. Profit = 50 + 70 = 120.
```

**Example 2:**
```
Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
Output: 150
Explanation: The subset chosen is the first, fourth and fifth job. Profit = 20 + 70 + 60 = 150.
```

## Constraints

- `1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4`
- `1 <= startTime[i] < endTime[i] <= 10^9`
- `1 <= profit[i] <= 10^4`

## Approach Hints

1. **Sort by end time + DP + Binary Search:** Sort jobs by end time. `dp[i]` = max profit considering first `i` jobs. For each job, binary search for the latest non-overlapping job.
2. **Key recurrence:** `dp[i] = max(dp[i-1], profit[i] + dp[last_non_overlapping])`.
3. **Alternative:** Use a TreeMap / sorted structure for efficient lookups.

## Related Problems

- [1751. Maximum Number of Events That Can Be Attended II](https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/)
- [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)
- [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

## Google Follow-ups

- What if each job has a priority and you want to maximize priority-weighted profit?
- How would you solve this if jobs can be partially completed?
- Can you parallelize this across multiple workers?
