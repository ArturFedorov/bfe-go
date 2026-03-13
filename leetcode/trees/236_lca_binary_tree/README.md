# 236. Lowest Common Ancestor of a Binary Tree

**Difficulty:** Medium

**Topics:** Tree, Depth-First Search, Binary Tree

## Description

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow a node to be a descendant of itself)."

## Examples

### Example 1
```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
```

### Example 2
```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself.
```

## Constraints

- The number of nodes in the tree is in the range `[2, 10^5]`.
- `-10^9 <= Node.val <= 10^9`
- All `Node.val` are unique.
- `p != q`
- `p` and `q` will exist in the tree.

## Approach Hints

1. **Recursive DFS:** If current node is p or q, return it. Recurse left and right. If both return non-nil, current node is the LCA.
2. **Parent pointers:** Build a parent map, then trace ancestors of p and q to find the first common one.
3. **Iterative with stack:** Use iterative post-order traversal to find both nodes and track ancestors.

## Related Problems

- 235. Lowest Common Ancestor of a Binary Search Tree
- 1644. Lowest Common Ancestor of a Binary Tree II
- 1676. Lowest Common Ancestor of a Binary Tree IV

## Google Follow-ups

- What if p or q might not exist in the tree?
- How would you find LCA with parent pointers available?
- Can you find the LCA of multiple nodes efficiently?
