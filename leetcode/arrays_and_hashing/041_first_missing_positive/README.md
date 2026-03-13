# 41. First Missing Positive

**Difficulty:** Hard

**Topics:** Array, Hash Table

## Description

Given an unsorted integer array `nums`, return the smallest missing positive integer.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

## Examples

### Example 1
```
Input: nums = [1,2,0]
Output: 3
Explanation: The numbers in the range [1,2] are all in the array.
```

### Example 2
```
Input: nums = [3,4,-1,1]
Output: 2
Explanation: 1 is in the array but 2 is missing.
```

### Example 3
```
Input: nums = [7,8,9,11,12]
Output: 1
Explanation: The smallest positive integer 1 is missing.
```

## Constraints

- `1 <= nums.length <= 10^5`
- `-2^31 <= nums[i] <= 2^31 - 1`

## Approach Hints

1. Use the array itself as a hash map by placing each number `nums[i]` at index `nums[i] - 1` (cyclic sort).
2. After rearranging, the first index `i` where `nums[i] != i + 1` gives the answer `i + 1`.
3. Ignore numbers that are out of range `[1, n]` since the answer must be in `[1, n+1]`.

## Related Problems

- [268. Missing Number](https://leetcode.com/problems/missing-number/)
- [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)
- [448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)

## What a Google Interviewer Would Ask Next

- **Why is the answer guaranteed to be in [1, n+1]?** With n slots, at most n distinct positive integers from 1..n can be present; so the first missing must be at most n+1.
- **Can you solve it with negative marking instead of swapping?** Yes—first replace non-positives with n+1, then use sign of nums[abs(val)-1] as a presence flag.
- **What if the input is a stream and you can't modify it?** You'd need O(n) space (e.g., a bit vector), since without random access or mutation O(1) space is impossible.
