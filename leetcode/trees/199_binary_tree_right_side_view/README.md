# 199. Binary Tree Right Side View

**Difficulty:** Medium

**Topics:** Tree, Depth-First Search, Breadth-First Search, Binary Tree

## Description

Given the `root` of a binary tree, imagine yourself standing on the **right side** of it, return the values of the nodes you can see ordered from top to bottom.

## Examples

### Example 1
```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```

### Example 2
```
Input: root = [1,null,3]
Output: [1,3]
```

### Example 3
```
Input: root = []
Output: []
```

## Constraints

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Approach Hints

1. **BFS level order:** Process each level, the last node in each level is visible from the right.
2. **DFS right-first:** Visit right child before left. The first node visited at each depth is the rightmost.
3. **DFS with depth tracking:** Track the current depth and add the node value if it's the first node at that depth.

## Related Problems

- 102. Binary Tree Level Order Traversal
- 116. Populating Next Right Pointers in Each Node
- 545. Boundary of Binary Tree

## Google Follow-ups

- How would you return the left side view instead?
- What if you need both left and right side views simultaneously?
- How would you handle this for a very deep tree to avoid stack overflow?
