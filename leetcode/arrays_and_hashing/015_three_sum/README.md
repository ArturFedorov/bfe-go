# 15. 3Sum

**Difficulty:** Medium

**Topics:** Array, Two Pointers, Sorting

## Description

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

## Examples

### Example 1
```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
```

### Example 2
```
Input: nums = [0,1,1]
Output: []
```

### Example 3
```
Input: nums = [0,0,0]
Output: [[0,0,0]]
```

## Constraints

- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Approach Hints

1. Sort the array first, then for each element use two pointers on the remaining subarray to find pairs that sum to the negation of the current element.
2. Skip duplicate values for the outer loop element and for both pointers to avoid duplicate triplets.
3. This reduces the problem from O(n^3) brute force to O(n^2) with the sort + two-pointer technique.

## Related Problems

- [1. Two Sum](https://leetcode.com/problems/two-sum/)
- [16. 3Sum Closest](https://leetcode.com/problems/3sum-closest/)
- [18. 4Sum](https://leetcode.com/problems/4sum/)

## What a Google Interviewer Would Ask Next

- **Can you do this in less than O(n^2)?** No, because in the worst case the output size itself can be O(n^2).
- **How would you handle 4Sum or kSum generically?** Recursively reduce kSum to 2Sum using sorting and two pointers, adding one loop per additional k.
- **What if the array is already sorted?** You skip the sort step but the algorithm is otherwise identical, still O(n^2).
