# 857. Minimum Cost to Hire K Workers

**Difficulty:** Hard

**Topics:** Array, Greedy, Sorting, Heap (Priority Queue)

## Description

There are `n` workers. You are given two integer arrays `quality` and `wage` where `quality[i]` is the quality of the `i`th worker and `wage[i]` is the minimum wage expectation for the `i`th worker.

We want to hire exactly `k` workers to form a paid group. When hiring a group of `k` workers, we must pay them according to the following rules:

1. Every worker in the paid group must be paid at least their minimum wage expectation.
2. In the group, each worker's pay is directly proportional to their quality relative to other workers in the group.

Given the integer `k`, return the least amount of money needed to form a paid group satisfying the above conditions. Answers within `10^-5` of the actual answer will be accepted.

## Examples

**Example 1:**
```
Input: quality = [10,20,5], wage = [70,50,30], k = 2
Output: 105.00000
Explanation: We pay 70 to worker 0 and 35 to worker 2.
```

**Example 2:**
```
Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
Output: 30.66667
```

## Constraints

- `n == quality.length == wage.length`
- `1 <= k <= n <= 10^4`
- `1 <= quality[i], wage[i] <= 10^4`

## Approach Hints

1. **Key insight:** The ratio wage[i]/quality[i] determines the rate. For a group, the rate must be the maximum ratio in the group, so cost = rate * totalQuality.
2. **Sort by ratio:** Sort workers by wage/quality ratio. Iterate through, maintaining a max-heap of qualities of size k. Replace the highest quality worker if adding a cheaper-rate worker reduces total cost.
3. **Time:** O(n log n + n log k). **Space:** O(n).

## Related Problems

- [347. Top K Frequent Elements](../347_top_k_frequent/)
- [502. IPO](https://leetcode.com/problems/ipo/)

## Google Follow-ups

- What if you can hire between k1 and k2 workers (a range)?
- What if workers have different availability times?
- How would you handle this if workers can be part-time (fractional quality)?
