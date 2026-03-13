# 34. Find First and Last Position of Element in Sorted Array

**Difficulty:** Medium

**Topics:** Array, Binary Search

## Description

Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

## Examples

**Example 1:**
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

**Example 2:**
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

**Example 3:**
```
Input: nums = [], target = 0
Output: [-1,-1]
```

## Constraints

- `0 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`
- `nums` is a non-decreasing array.
- `-10^9 <= target <= 10^9`

## Approach Hints

1. **Two binary searches:** One to find the leftmost (first) occurrence, one to find the rightmost (last) occurrence.
2. **Left search:** Find the smallest index where `nums[i] >= target`, verify it equals target.
3. **Right search:** Find the largest index where `nums[i] <= target`, or use `leftSearch(target+1) - 1`.

## Related Problems

- [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)
- [278. First Bad Version](https://leetcode.com/problems/first-bad-version/)
- [2089. Find Target Indices After Sorting](https://leetcode.com/problems/find-target-indices-after-sorting-array/)

## Google Follow-ups

- How would you count the total occurrences of a target in O(log n)?
- What if the array is sorted in non-increasing order?
- How would you extend this to a 2D sorted matrix?
