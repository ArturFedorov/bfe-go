# 39. Combination Sum

**Difficulty:** Medium

**Topics:** Array, Backtracking

## Description

Given an array of distinct integers `candidates` and a target integer `target`, return a list of all unique combinations of `candidates` where the chosen numbers sum to `target`. You may return the combinations in any order.

The same number may be chosen from `candidates` an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

## Examples

**Example 1:**
```
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
```

**Example 2:**
```
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
```

**Example 3:**
```
Input: candidates = [2], target = 1
Output: []
```

## Constraints

- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- All elements of `candidates` are distinct.
- `1 <= target <= 40`

## Approach Hints

1. **Backtracking with index:** At each step, try adding each candidate (from current index onward to avoid duplicates) and recurse with reduced target.
2. **Sort first:** Sorting candidates allows early termination when remaining candidates exceed target.
3. **Time:** O(n^(T/M)) where T is target and M is minimum candidate. **Space:** O(T/M) for recursion depth.

## Related Problems

- [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)
- [46. Permutations](../046_permutations/)
- [77. Combinations](https://leetcode.com/problems/combinations/)

## Google Follow-ups

- What if candidates can contain duplicates? How do you avoid duplicate combinations?
- What if you need to return the combination closest to the target if no exact match exists?
- How would you parallelize the backtracking search?
