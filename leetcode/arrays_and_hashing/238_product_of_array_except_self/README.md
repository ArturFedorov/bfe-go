# 238. Product of Array Except Self

**Difficulty:** Medium

**Topics:** Array, Prefix Sum

## Description

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

## Examples

### Example 1
```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

### Example 2
```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
```

## Constraints

- `2 <= nums.length <= 10^5`
- `-30 <= nums[i] <= 30`
- The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

## Follow-up

Can you solve it in O(1) extra space complexity? (The output array does not count as extra space.)

## Approach Hints

1. **Brute Force:** For each element, multiply all other elements. O(n²) time.
2. **Prefix and Suffix Arrays:** Build left-product and right-product arrays, then multiply them. O(n) time, O(n) space.
3. **Single Output Array:** Use the output array for left products, then sweep right-to-left with a running product. O(n) time, O(1) extra space.

## Related Problems

- [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)
- [152. Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)
- [265. Paint House II](https://leetcode.com/problems/paint-house-ii/)

## What a Google Interviewer Would Ask Next

- **What if division were allowed?** Compute total product, then divide by each element. But you must handle zeros carefully (one zero vs multiple zeros).
- **How do you handle zeros?** With one zero, only that position gets a non-zero result. With two or more zeros, all results are zero.
- **What about integer overflow?** Use long/int64, or apply modular arithmetic if working under a modulus. The constraint here guarantees 32-bit safety.
- **How would you parallelize this?** Compute prefix products in parallel using a parallel prefix sum (scan) algorithm, then combine with suffix products.
