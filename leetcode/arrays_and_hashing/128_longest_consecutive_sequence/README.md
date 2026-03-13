# 128. Longest Consecutive Sequence

**Difficulty:** Medium

**Topics:** Array, Hash Table, Union Find

## Description

Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

## Examples

### Example 1
```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

### Example 2
```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

## Constraints

- `0 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Approach Hints

1. Insert all numbers into a hash set for O(1) lookups.
2. For each number, only start counting a sequence if `num - 1` is NOT in the set (this ensures you only start from the beginning of a sequence).
3. From each sequence start, count consecutive numbers upward and track the maximum length.

## Related Problems

- [298. Binary Tree Longest Consecutive Sequence](https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/)
- [2177. Find Three Consecutive Integers That Sum to a Given Number](https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/)

## What a Google Interviewer Would Ask Next

- **Why is this O(n) even though there's a nested while loop?** Each element is visited at most twice (once in the outer loop, once in a while expansion), so total work is O(n).
- **Could you solve this with Union-Find?** Yes, union consecutive numbers and track component sizes, also O(n) with path compression.
- **What if the input is a stream?** Maintain a hash map of interval boundaries; when a new number arrives, merge with adjacent intervals in O(1).
