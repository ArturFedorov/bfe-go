# 23. Merge k Sorted Lists

**Difficulty:** Hard

**Topics:** Linked List, Divide and Conquer, Heap (Priority Queue), Merge Sort

## Description

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

## Examples

**Example 1:**

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
```

**Example 2:**

```
Input: lists = []
Output: []
```

**Example 3:**

```
Input: lists = [[]]
Output: []
```

## Constraints

- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` is sorted in ascending order.
- The sum of `lists[i].length` will not exceed `10^4`.

## What a Google Interviewer Would Ask Next

- **Heap vs divide-and-conquer approach trade-offs?** A min-heap approach gives O(N log k) time with O(k) space for the heap. Divide-and-conquer also achieves O(N log k) time but with O(log k) space for the recursion stack. The heap approach processes one node at a time (better for streaming), while divide-and-conquer has better cache locality due to sequential merges.
- **How would you handle streaming input?** If lists arrive over time, a min-heap approach is ideal since you can insert new list heads dynamically. You could maintain a persistent heap and merge new lists as they arrive without restarting the entire computation.
- **What's the time complexity of each approach?** Brute force (collect all, sort): O(N log N). Merge one by one: O(kN). Merge with heap: O(N log k). Divide and conquer: O(N log k). Here N is the total number of nodes across all lists, and k is the number of lists.
- **How would you parallelize this?** Divide-and-conquer naturally parallelizes — assign pairs of lists to different threads/cores for merging at each level. With k lists, you get log(k) rounds where each round's merges are independent and can run concurrently, achieving near-linear speedup with up to k/2 processors.
- **What if lists are on different machines (external merge sort)?** Use a distributed min-heap where each machine sends its current minimum. The coordinator maintains a k-way merge with network-buffered reads to amortize latency. This is the core of external merge sort used in MapReduce — each machine sorts locally, then a merge phase combines sorted runs using bounded memory buffers.
