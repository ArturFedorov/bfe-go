# 543. Diameter of Binary Tree

**Difficulty:** Easy

**Topics:** Tree, Depth-First Search, Binary Tree

## Description

Given the `root` of a binary tree, return the length of the **diameter** of the tree.

The **diameter** of a binary tree is the **length** of the longest path between any two nodes in a tree. This path may or may not pass through the `root`.

The **length** of a path between two nodes is represented by the number of edges between them.

## Examples

### Example 1
```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
```

### Example 2
```
Input: root = [1,2]
Output: 1
```

## Constraints

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-100 <= Node.val <= 100`

## Approach Hints

1. **DFS with height:** At each node, the diameter through it is leftHeight + rightHeight. Track the global maximum.
2. **Post-order traversal:** Compute heights bottom-up and update diameter at each node.
3. **Similar to 124:** Same pattern of computing a local value while tracking a global optimum.

## Related Problems

- 104. Maximum Depth of Binary Tree
- 124. Binary Tree Maximum Path Sum
- 687. Longest Univalue Path

## Google Follow-ups

- What if you need to return the actual path, not just the length?
- How would you find the diameter of an n-ary tree?
- Can you compute the diameter iteratively?
