# 951. Flip Equivalent Binary Trees

**Difficulty:** Medium

**Topics:** Tree, Depth-First Search, Binary Tree

## Description

For a binary tree `T`, we can define a **flip operation** as follows: choose any node, and swap the left and right child subtrees.

A binary tree `X` is **flip equivalent** to a binary tree `Y` if and only if we can make `X` equal to `Y` after some number of flip operations.

Given the roots of two binary trees `root1` and `root2`, return `true` if the two trees are flip equivalent or `false` otherwise.

## Examples

### Example 1
```
Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
Output: true
Explanation: We flipped at nodes with values 1, 3, and 5.
```

### Example 2
```
Input: root1 = [], root2 = []
Output: true
```

### Example 3
```
Input: root1 = [], root2 = [1]
Output: false
```

## Constraints

- The number of nodes in each tree is in the range `[0, 100]`.
- Each tree will have unique node values in the range `[0, 99]`.

## Approach Hints

1. **Recursive:** Two trees are flip equivalent if roots match and either (left1==left2 AND right1==right2) OR (left1==right2 AND right1==left2).
2. **Canonical form:** Convert both trees to a canonical form (e.g., smaller child always on left), then compare.
3. **BFS comparison:** Compare level by level, considering both orderings at each node.

## Related Problems

- 100. Same Tree
- 226. Invert Binary Tree
- 572. Subtree of Another Tree

## Google Follow-ups

- What is the minimum number of flips needed to make the trees equivalent?
- How would you extend this to n-ary trees?
- Can you determine flip equivalence without recursion using O(1) space?
