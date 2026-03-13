# 297. Serialize and Deserialize Binary Tree

**Difficulty:** Hard

**Topics:** Tree, Depth-First Search, Breadth-First Search, Design, String, Binary Tree

## Description

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

## Examples

### Example 1
```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```

### Example 2
```
Input: root = []
Output: []
```

## Constraints

- The number of nodes in the tree is in the range `[0, 10^4]`.
- `-1000 <= Node.val <= 1000`

## Approach Hints

1. **Pre-order DFS:** Serialize using pre-order traversal with "null" markers. Deserialize by reading tokens in the same order.
2. **BFS level-order:** Serialize level by level, use "null" for missing children. Deserialize using a queue.
3. **Key insight:** You need null markers to uniquely reconstruct the tree from a single traversal.

## Related Problems

- 449. Serialize and Deserialize BST
- 428. Serialize and Deserialize N-ary Tree
- 652. Find Duplicate Subtrees

## Google Follow-ups

- How would you optimize the serialized format for space?
- Can you serialize/deserialize a binary tree using only two traversals (no null markers)?
- How would you handle very large trees that don't fit in memory?
