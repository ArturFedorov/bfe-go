# 399. Evaluate Division

**Difficulty:** Medium

**Topics:** Array, Depth-First Search, Breadth-First Search, Union Find, Graph, Shortest Path

## Description

You are given an array of variable pairs `equations` and an array of real numbers `values`, where `equations[i] = [Ai, Bi]` and `values[i]` represent the equation `Ai / Bi = values[i]`. Each `Ai` or `Bi` is a string that represents a single variable.

You are also given some `queries`, where `queries[j] = [Cj, Dj]` represents the `jth` query where you must find the answer for `Cj / Dj = ?`.

Return the answers to all queries. If a single answer cannot be determined, return `-1.0`.

## Examples

### Example 1
```
Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
```

## Constraints

- `1 <= equations.length <= 20`
- `equations[i].length == 2`
- `1 <= Ai.length, Bi.length <= 5`
- `values.length == equations.length`
- `0.0 < values[i] <= 20.0`
- `1 <= queries.length <= 20`
- `queries[j].length == 2`
- `1 <= Cj.length, Dj.length <= 5`
- `Ai, Bi, Cj, Dj` consist of lower case English letters and digits.

## Approach Hints

1. **Graph DFS/BFS:** Build a weighted directed graph. For a/b=k, add edges a->b with weight k and b->a with weight 1/k. Answer queries by finding paths.
2. **Union Find with weights:** Use weighted union find where each node stores its ratio to the root.
3. **Floyd-Warshall:** Build a matrix of all-pairs ratios using transitive closure.

## Related Problems

- 990. Satisfiability of Equality Equations
- 721. Accounts Merge
- 952. Largest Component Size by Common Factor

## Google Follow-ups

- How would you handle updates (changing equation values)?
- What if equations can be contradictory? How do you detect inconsistencies?
- How would you scale this to millions of variables?
