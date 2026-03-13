# 206. Reverse Linked List

**Difficulty:** Easy

**Topics:** Linked List, Recursion

## Description

Given the head of a singly linked list, reverse the list, and return the reversed list.

## Examples

**Example 1:**

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

**Example 2:**

```
Input: head = [1,2]
Output: [2,1]
```

**Example 3:**

```
Input: head = []
Output: []
```

## Constraints

- The number of nodes in the list is in the range `[0, 5000]`.
- `-5000 <= Node.val <= 5000`

**Follow up:** A linked list can be reversed either iteratively or recursively. Could you implement both?

## Approach Hints

**Iterative:** Use three pointers (`prev`, `curr`, `next`). Walk through the list, reversing each node's `Next` pointer to point to the previous node. Time O(n), Space O(1).

**Recursive:** Reverse the rest of the list first, then fix the pointers. The base case is when the list is empty or has one node. Time O(n), Space O(n) due to call stack.

## Related Problems

- [92. Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)
- [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

## What a Google Interviewer Would Ask Next

- **Reverse in groups of k:** "Now reverse the nodes in groups of k. If the remaining nodes are fewer than k, leave them as-is." (LeetCode 25)
- **Detect cycle first:** "What if the list might have a cycle? How would you detect it before attempting to reverse?"
- **Doubly linked list:** "How would your solution change if each node also had a `Prev` pointer?"
- **Circular linked list:** "What if the list is circular — the tail points back to the head? How would you reverse it and maintain the circular structure?"
- **Space complexity:** "Your iterative solution is already O(1) space. Can you make the recursive solution use O(1) space as well?" (Hint: tail recursion, though Go doesn't optimize for it.)
