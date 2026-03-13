# 124. Binary Tree Maximum Path Sum

**Difficulty:** Hard

**Topics:** Tree, Depth-First Search, Dynamic Programming, Binary Tree

## Description

A **path** in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence **at most once**. Note that the path does not need to pass through the root.

The **path sum** of a path is the sum of the node's values in the path.

Given the `root` of a binary tree, return the maximum **path sum** of any **non-empty** path.

## Examples

### Example 1
```
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
```

### Example 2
```
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
```

## Constraints

- The number of nodes in the tree is in the range `[1, 3 * 10^4]`.
- `-1000 <= Node.val <= 1000`

## Approach Hints

1. **Post-order DFS:** At each node, compute the max gain from left and right subtrees. Update global max considering the path through the current node.
2. **Key insight:** A node can either be part of a path going up to its parent OR be the "turning point" where left + node + right forms the max path.
3. **Negative subtrees:** If a subtree contributes a negative sum, ignore it (use 0 instead).

## Related Problems

- 112. Path Sum
- 129. Sum Root to Leaf Numbers
- 687. Longest Univalue Path

## Google Follow-ups

- What if you need to return the actual path, not just the sum?
- How would you handle this for an n-ary tree?
- Can you find the top-K maximum path sums efficiently?
