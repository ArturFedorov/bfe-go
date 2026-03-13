# 207. Course Schedule

**Difficulty:** Medium

**Topics:** Depth-First Search, Breadth-First Search, Graph, Topological Sort

## Description

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi` first if you want to take course `ai`.

For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return `true` if you can finish all courses. Otherwise, return `false`.

## Examples

### Example 1
```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are 2 courses to take. To take course 1 you should have finished course 0. So it is possible.
```

### Example 2
```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are 2 courses to take. To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
```

## Constraints

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- All the pairs `prerequisites[i]` are **unique**.

## Approach Hints

1. **Topological sort (Kahn's/BFS):** Build in-degree array. Start with nodes of in-degree 0. If all nodes are processed, no cycle exists.
2. **DFS cycle detection:** Use three states (unvisited, visiting, visited). A back edge to a "visiting" node means a cycle.
3. **Key insight:** The problem reduces to detecting a cycle in a directed graph.

## Related Problems

- 210. Course Schedule II
- 269. Alien Dictionary
- 310. Minimum Height Trees

## Google Follow-ups

- What if courses have time durations and you want the minimum total time with parallel execution?
- How would you handle dynamic prerequisites that change over time?
- What if there are soft prerequisites (recommended but not required)?
