# 703. Kth Largest Element in a Stream

**Difficulty:** Easy

**Topics:** Tree, Design, Binary Search Tree, Heap (Priority Queue), Binary Tree, Data Stream

## Description

Design a class to find the `k`th largest element in a stream. Note that it is the `k`th largest element in the sorted order, not the `k`th distinct element.

Implement `KthLargest` class:

- `KthLargest(int k, int[] nums)` Initializes the object with the integer `k` and the stream of integers `nums`.
- `int add(int val)` Appends the integer `val` to the stream and returns the element representing the `k`th largest element in the stream.

## Examples

**Example 1:**

```
Input: ["KthLargest", "add", "add", "add", "add", "add"]
       [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output: [null, 4, 5, 5, 8, 8]
```

## Constraints

- `1 <= k <= 10^4`
- `0 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `-10^4 <= val <= 10^4`
- At most `10^4` calls will be made to `add`.
- It is guaranteed that there will be at least `k` elements in the array when you search for the `k`th element.

## Approach Hints

1. **Min-heap of size k:** Maintain a min-heap of size `k`. The root is always the kth largest. O(log k) per add.
2. **Sorted slice with binary search:** Keep a sorted slice and use binary search for insertion. O(n) per add due to shifting.
3. **Balanced BST with order statistics:** Augmented BST that tracks subtree sizes for O(log n) kth element queries.

## Related Problems

- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)
- [295. Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)

## Google Follow-ups

- How would you handle this if the stream is distributed across multiple machines?
- What if you need to support removing elements as well?
- How would you modify this to find the kth largest in a sliding window?
