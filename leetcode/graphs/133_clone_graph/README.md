# 133. Clone Graph

**Difficulty:** Medium

**Topics:** Hash Table, Depth-First Search, Breadth-First Search, Graph

## Description

Given a reference of a node in a **connected** undirected graph, return a **deep copy** (clone) of the graph.

Each node in the graph contains a value (`int`) and a list (`List[Node]`) of its neighbors.

## Examples

### Example 1
```
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
```

### Example 2
```
Input: adjList = [[]]
Output: [[]]
Explanation: The graph consists of a single node with no neighbors.
```

### Example 3
```
Input: adjList = []
Output: []
Explanation: The graph is empty (null).
```

## Constraints

- The number of nodes in the graph is in the range `[0, 100]`.
- `1 <= Node.val <= 100`
- `Node.val` is unique for each node.
- There are no repeated edges and no self-loops in the graph.
- The Graph is connected and all nodes can be visited starting from the given node.

## Approach Hints

1. **DFS with hash map:** Use a map from original node to cloned node. For each node, create its clone and recursively clone neighbors.
2. **BFS with hash map:** Use a queue for BFS traversal and a map to track cloned nodes.
3. **Key insight:** The map serves dual purpose: tracking visited nodes and mapping originals to clones.

## Related Problems

- 138. Copy List with Random Pointer
- 1485. Clone Binary Tree With Random Pointer
- 1490. Clone N-ary Tree

## Google Follow-ups

- How would you clone a graph with weighted edges?
- What if the graph is very large and distributed across multiple machines?
- How would you verify that a deep copy is correct?
