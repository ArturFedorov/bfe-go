# 98. Validate Binary Search Tree

**Difficulty:** Medium

**Topics:** Tree, Depth-First Search, Binary Search Tree, Binary Tree

## Description

Given the `root` of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

## Examples

### Example 1
```
Input: root = [2,1,3]
Output: true
```

### Example 2
```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

## Constraints

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-2^31 <= Node.val <= 2^31 - 1`

## Approach Hints

1. **Recursive with range:** Pass min/max bounds down the recursion. Each node must be within (min, max).
2. **In-order traversal:** An in-order traversal of a valid BST produces a strictly increasing sequence.
3. **Iterative with stack:** Use an explicit stack to simulate in-order traversal.

## Related Problems

- 94. Binary Tree Inorder Traversal
- 501. Find Mode in Binary Search Tree
- 700. Search in a Binary Search Tree

## Google Follow-ups

- What if the BST allows duplicate values on the left side?
- Can you solve it iteratively with O(1) space (Morris traversal)?
- How would you validate a BST stored as a sorted array?
