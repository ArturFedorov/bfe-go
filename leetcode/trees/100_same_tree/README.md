# 100. Same Tree

**Difficulty:** Easy

**Topics:** Tree, Depth-First Search, Breadth-First Search, Binary Tree

## Description

Given the roots of two binary trees `p` and `q`, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

## Examples

### Example 1
```
Input: p = [1,2,3], q = [1,2,3]
Output: true
```

### Example 2
```
Input: p = [1,2], q = [1,null,2]
Output: false
```

### Example 3
```
Input: p = [1,2,1], q = [1,1,2]
Output: false
```

## Constraints

- The number of nodes in both trees is in the range `[0, 100]`.
- `-10^4 <= Node.val <= 10^4`

## Approach Hints

1. **Recursive DFS:** Compare current nodes, then recurse on left and right children.
2. **Iterative BFS:** Use two queues and compare nodes level by level.
3. **Serialization:** Serialize both trees and compare strings (less efficient but valid).

## Related Problems

- 101. Symmetric Tree
- 572. Subtree of Another Tree
- 951. Flip Equivalent Binary Trees

## Google Follow-ups

- How would you compare two very large trees that don't fit in memory?
- Can you check if two trees are the same using only O(1) extra space?
- What if you need to compare trees stored across distributed systems?
