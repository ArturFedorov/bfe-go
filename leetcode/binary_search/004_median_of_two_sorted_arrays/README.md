# 4. Median of Two Sorted Arrays

**Difficulty:** Hard

**Topics:** Array, Binary Search, Divide and Conquer

## Description

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

## Examples

**Example 1:**
```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```

**Example 2:**
```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

## Constraints

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

## Approach Hints

1. **Binary search on shorter array:** Partition both arrays such that left halves combined equal right halves. Binary search the partition point on the shorter array.
2. **Key invariant:** `maxLeft1 <= minRight2` and `maxLeft2 <= minRight1`.
3. **Edge cases:** Handle empty arrays and odd/even total lengths carefully.

## Related Problems

- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)
- [295. Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)

## Google Follow-ups

- Can you generalize this to find the k-th element of two sorted arrays?
- What if there are more than two sorted arrays?
- How would you handle this in a distributed system where arrays are on different machines?
