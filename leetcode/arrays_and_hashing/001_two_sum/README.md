# 1. Two Sum

**Difficulty:** Easy

**Topics:** Array, Hash Table

## Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Examples

### Example 1
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

### Example 2
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

### Example 3
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

## Constraints

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- Only one valid answer exists.

## Approach Hints

1. **Brute Force:** Check every pair of numbers. O(n²) time, O(1) space.
2. **Hash Map:** For each number, check if `target - num` exists in the map. O(n) time, O(n) space.
3. **Sorting + Two Pointers:** Sort and use two pointers, but you need to track original indices. O(n log n) time, O(n) space.

## Related Problems

- [15. 3Sum](https://leetcode.com/problems/3sum/)
- [18. 4Sum](https://leetcode.com/problems/4sum/)
- [167. Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)
- [170. Two Sum III - Data Structure Design](https://leetcode.com/problems/two-sum-iii-data-structure-design/)

## What a Google Interviewer Would Ask Next

- **What if the array is sorted?** Use two pointers from both ends — O(n) time, O(1) space.
- **How would you extend this to 3Sum or 4Sum?** Sort the array and reduce to Two Sum with nested loops.
- **What if there are multiple valid pairs?** Return all pairs — use a hash map and handle duplicates carefully.
- **What if the array is very large and doesn't fit in memory?** Distribute data across machines, use external sort + two pointers, or partition by hash and solve per-partition.
