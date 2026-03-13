# 153. Find Minimum in Rotated Sorted Array

**Difficulty:** Medium

**Topics:** Array, Binary Search

## Description

Suppose an array of length `n` sorted in ascending order is **rotated** between `1` and `n` times. For example, the array `nums = [0,1,2,4,5,6,7]` might become:

- `[4,5,6,7,0,1,2]` if it was rotated 4 times.
- `[0,1,2,4,5,6,7]` if it was rotated 7 times.

Notice that **rotating** an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` of **unique** elements, return the minimum element of this array.

You must write an algorithm that runs in `O(log n)` time.

## Examples

**Example 1:**
```
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
```

**Example 2:**
```
Input: nums = [4,5,6,7,0,1,2]
Output: 0
```

**Example 3:**
```
Input: nums = [11,13,15,17]
Output: 11
Explanation: The array was not rotated (or rotated n times).
```

## Constraints

- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- All the integers of `nums` are **unique**.
- `nums` is sorted and rotated between `1` and `n` times.

## Approach Hints

1. **Binary search:** Compare `nums[mid]` with `nums[right]`. If `nums[mid] > nums[right]`, min is in the right half; otherwise, it's in the left half (including mid).
2. **Key insight:** The minimum is the only element where `nums[i] < nums[i-1]` (or it's the first element if not rotated).
3. **Invariant:** Maintain the search space such that the minimum is always within `[left, right]`.

## Related Problems

- [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)
- [154. Find Minimum in Rotated Sorted Array II](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)

## Google Follow-ups

- What if duplicates are allowed (problem 154)?
- Can you also find the number of rotations?
- How would you handle this for a circularly sorted linked list?
