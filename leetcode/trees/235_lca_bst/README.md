# 235. Lowest Common Ancestor of a Binary Search Tree

**Difficulty:** Medium

**Topics:** Tree, Depth-First Search, Binary Search Tree, Binary Tree

## Description

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow a node to be a descendant of itself)."

## Examples

### Example 1
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
```

### Example 2
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself.
```

## Constraints

- The number of nodes in the tree is in the range `[2, 10^5]`.
- `-10^9 <= Node.val <= 10^9`
- All `Node.val` are unique.
- `p != q`
- `p` and `q` will exist in the BST.

## Approach Hints

1. **BST property:** If both p and q are smaller than root, LCA is in left subtree. If both are larger, LCA is in right subtree. Otherwise, root is the LCA.
2. **Iterative:** Follow the BST property without recursion for O(1) space.
3. **Compare with generic LCA (236):** BST property makes this more efficient than the general binary tree case.

## Related Problems

- 236. Lowest Common Ancestor of a Binary Tree
- 1644. Lowest Common Ancestor of a Binary Tree II
- 1650. Lowest Common Ancestor of a Binary Tree III

## Google Follow-ups

- What if the BST is very deep and you need to avoid stack overflow?
- How would you find the LCA of k nodes in a BST?
- What if nodes p and q may not exist in the tree?
