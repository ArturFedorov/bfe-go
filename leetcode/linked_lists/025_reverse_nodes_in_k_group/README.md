# 25. Reverse Nodes in k-Group

**Difficulty:** Hard
**Topics:** Linked List, Recursion

## Description

Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return the modified list.

`k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

## Examples

### Example 1

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

### Example 2

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

## Constraints

- The number of nodes in the list is `n`.
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

## Follow-up

Can you solve the problem in `O(1)` extra memory space?

## What a Google Interviewer Would Ask Next

1. **Iterative vs Recursive:** Can you implement this both iteratively and recursively? What are the trade-offs in terms of space complexity? The recursive approach uses O(n/k) stack space — can you eliminate that?

2. **Reverse the remainder too:** What if instead of leaving the final group as-is, we should also reverse it even if it has fewer than `k` nodes? How would you modify your solution?

3. **Doubly linked list variant:** How would your approach change if the input were a doubly linked list? Would it simplify or complicate the reversal logic?

4. **Very large k:** If `k` could be much larger than `n`, how would you handle it efficiently without traversing the list multiple times to count nodes?

5. **Relation to LC 206 (Reverse Linked List):** How does this problem build on reversing an entire linked list? Can you decompose this problem into repeated applications of that simpler operation?
