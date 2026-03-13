# 295. Find Median from Data Stream

**Difficulty:** Hard

**Topics:** Two Heaps, Design, Sorting, Data Stream

## Description

The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

Implement the `MedianFinder` class:

- `MedianFinder()` initializes the `MedianFinder` object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far.

## Examples

**Example 1:**

```
Input: ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
       [[], [1], [2], [], [3], []]
Output: [null, null, null, 1.5, null, 2.0]
```

## Constraints

- `-10^5 <= num <= 10^5`
- There will be at least one element in the data structure before calling `findMedian`.
- At most `5 * 10^4` calls will be made to `addNum` and `findMedian`.

## Approach Hints

1. **Two heaps:** Use a max-heap for the lower half and a min-heap for the upper half. Keep them balanced.
2. **Sorted list with binary search insertion:** Maintain a sorted slice; use `sort.SearchInts` for O(log n) lookup but O(n) insertion.
3. **Balanced BST / order-statistic tree:** Allows O(log n) insertion and median lookup.

## Related Problems

- [480. Sliding Window Median](https://leetcode.com/problems/sliding-window-median/)
- [502. IPO](https://leetcode.com/problems/ipo/)

## Google Follow-ups

- How would you handle it if all integers are in the range [0, 100]? (Bucket counting)
- How would you handle it if 99% of all integers are in the range [0, 100]? (Bucket + overflow)
- How would you design this for a distributed system with multiple data streams?
