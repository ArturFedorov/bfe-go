# 21. Merge Two Sorted Lists

**Difficulty:** Easy

**Topics:** Linked List, Recursion

---

## Description

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return _the head of the merged linked list_.

---

## Examples

### Example 1

```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

### Example 2

```
Input: list1 = [], list2 = []
Output: []
```

### Example 3

```
Input: list1 = [], list2 = [0]
Output: [0]
```

---

## Constraints

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.

---

## Approach Hints

<details>
<summary>Hint 1</summary>
Use a dummy head node to simplify edge cases when building the merged list.
</details>

<details>
<summary>Hint 2</summary>
Compare the current nodes of both lists and attach the smaller one to the result. Advance that list's pointer.
</details>

<details>
<summary>Hint 3</summary>
When one list is exhausted, append the remainder of the other list directly.
</details>

---

## Related Problems

- [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) (Hard)
- [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) (Easy)
- [148. Sort List](https://leetcode.com/problems/sort-list/) (Medium)

---

### What a Google Interviewer Would Ask Next
```
1. "How would you merge k sorted lists instead of just two?" (LC 23)
   → Use a min-heap (priority queue) of size k to always pick the smallest head.
   → Alternatively, divide-and-conquer: pair up lists and merge recursively.
   → Time: O(N log k) where N is total nodes across all lists.

2. "What are the trade-offs between iterative and recursive approaches?"
   → Iterative: O(1) space, no risk of stack overflow on long lists.
   → Recursive: cleaner code, but O(n+m) stack space — problematic for large inputs.
   → In production, iterative is almost always preferred.

3. "What if the lists are doubly linked?"
   → Same merge logic, but also maintain prev pointers during splicing.
   → Extra bookkeeping but same O(n+m) time complexity.

4. "Should you merge in-place or create new nodes?"
   → In-place (rewiring pointers): O(1) extra space, but mutates input.
   → Creating new nodes: O(n+m) extra space, but preserves original lists.
   → Clarify with interviewer which is expected.

5. "How does this relate to merge sort?"
   → This is the 'merge' step of merge sort applied to linked lists.
   → Merge sort on a linked list uses this as a subroutine with
     top-down (recursive split) or bottom-up (iterative doubling) decomposition.
   → Linked list merge sort is O(n log n) time, O(1) extra space (bottom-up).
```
