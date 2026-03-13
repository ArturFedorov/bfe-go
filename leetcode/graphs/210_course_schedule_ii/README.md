# 210. Course Schedule II

**Difficulty:** Medium

**Topics:** Depth-First Search, Breadth-First Search, Graph, Topological Sort

## Description

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi` first if you want to take course `ai`.

For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return **any** of them. If it is impossible to finish all courses, return **an empty array**.

## Examples

### Example 1
```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
```

### Example 2
```
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3] or [0,1,2,3]
Explanation: There are 4 courses to take. Both orderings are correct.
```

### Example 3
```
Input: numCourses = 1, prerequisites = []
Output: [0]
```

## Constraints

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `ai != bi`
- All the pairs `[ai, bi]` are **distinct**.

## Approach Hints

1. **Kahn's algorithm (BFS):** Build in-degree array, process nodes with in-degree 0, output them in order.
2. **DFS with post-order:** Run DFS, add nodes to result in reverse post-order (when all descendants are processed).
3. **Cycle detection:** If the result doesn't include all courses, there's a cycle — return empty array.

## Related Problems

- 207. Course Schedule
- 269. Alien Dictionary
- 1136. Parallel Courses

## Google Follow-ups

- What if you want to minimize the number of semesters (parallel execution)?
- How would you handle weighted prerequisites (some take longer)?
- Return all valid topological orderings.
