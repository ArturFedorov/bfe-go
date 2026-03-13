# 2. Add Two Numbers

**Difficulty:** Medium

**Topics:** Linked List, Math, Recursion

## Description

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

## Examples

**Example 1:**

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807
```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

## Constraints

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

## What a Google Interviewer Would Ask Next

- **What if the digits are stored in forward order instead of reverse?** This is LeetCode 445 (Add Two Numbers II). You would need to either reverse the lists first or use a stack-based approach.
- **What about multiplying two numbers represented as linked lists?** Similar traversal logic but requires handling partial products and accumulating results, significantly more complex carry management.
- **How would you handle very large numbers (BigInt)?** The linked list representation already handles arbitrarily large numbers naturally, since each node stores a single digit. Discuss trade-offs vs array-based or string-based BigInt implementations.
- **What if one list is much longer than the other?** Optimize by detecting the shorter list and only iterating with carry propagation through the remainder of the longer list once the shorter one is exhausted.
