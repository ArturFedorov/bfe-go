# 141. Linked List Cycle

**Difficulty:** Easy

**Topics:** Linked List, Hash Table, Two Pointers

## Description

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter.**

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

## Examples

**Example 1:**

```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```

**Example 2:**

```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
```

**Example 3:**

```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```

## Constraints

- The number of nodes in the list is in the range `[0, 10^4]`.
- `-10^5 <= Node.val <= 10^5`
- `pos` is `-1` or a valid index in the linked list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

## Approach Hints

**Hash Set:** Walk the list and store visited nodes in a set. If you see a node you've already visited, there's a cycle. Time O(n), Space O(n).

**Floyd's Tortoise and Hare:** Use two pointers — slow moves one step, fast moves two steps. If they meet, there's a cycle. If fast reaches nil, there's no cycle. Time O(n), Space O(1).

## Related Problems

- [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)
- [202. Happy Number](https://leetcode.com/problems/happy-number/)
- [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

## What a Google Interviewer Would Ask Next

- **Find the cycle start:** "Can you return the node where the cycle begins?" (LeetCode 142)
- **Cycle length:** "Once you've detected a cycle, how would you find its length?"
- **Remove the cycle:** "How would you break the cycle so the list becomes acyclic?"
- **Space trade-off:** "Your hash set approach uses O(n) space. Can you prove why Floyd's algorithm is guaranteed to detect a cycle in O(n) time?"
- **Corrupted list:** "What if the list is doubly linked — how would cycle detection change?"
