# 347. Top K Frequent Elements

**Difficulty:** Medium

**Topics:** Array, Hash Table, Divide and Conquer, Sorting, Heap (Priority Queue), Bucket Sort, Counting, Quickselect

## Description

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

## Examples

**Example 1:**
```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**
```
Input: nums = [1], k = 1
Output: [1]
```

## Constraints

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is guaranteed that the answer is unique.

## Approach Hints

1. **Min-heap of size k:** Count frequencies with a hash map, then maintain a min-heap of size k.
2. **Bucket sort:** Use frequency as index into buckets. Iterate from highest frequency bucket.
3. **Quickselect:** Partition by frequency to find the k-th most frequent in O(n) average.
4. **Time:** O(n log k) for heap, O(n) for bucket sort. **Space:** O(n).

## Related Problems

- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)
- [692. Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/)
- [621. Task Scheduler](../621_task_scheduler/)

## Google Follow-ups

- Can you solve it in O(n) time?
- What if the data is streaming? How would you maintain top-k frequent elements?
- How would you handle ties in frequency?
