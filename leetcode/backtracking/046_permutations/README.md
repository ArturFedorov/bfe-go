# 46. Permutations

**Difficulty:** Medium

**Topics:** Array, Backtracking

## Description

Given an array `nums` of distinct integers, return all the possible permutations. You can return the answer in any order.

## Examples

**Example 1:**
```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**Example 2:**
```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

**Example 3:**
```
Input: nums = [1]
Output: [[1]]
```

## Constraints

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- All the integers of `nums` are unique.

## Approach Hints

1. **Backtracking with swap:** Swap each element into the current position and recurse on the remaining.
2. **Backtracking with visited:** Use a boolean array to track which elements have been used.
3. **Time:** O(n! * n). **Space:** O(n) for recursion depth.

## Related Problems

- [47. Permutations II](https://leetcode.com/problems/permutations-ii/)
- [39. Combination Sum](../039_combination_sum/)
- [51. N-Queens](../051_n_queens/)

## Google Follow-ups

- How would you generate the k-th permutation directly without generating all permutations?
- What if the array contains duplicates?
- How would you generate permutations in lexicographic order?
