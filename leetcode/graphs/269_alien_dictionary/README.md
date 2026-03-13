# 269. Alien Dictionary

**Difficulty:** Hard

**Topics:** Array, String, Depth-First Search, Breadth-First Search, Graph, Topological Sort

## Description

There is a new alien language that uses the English alphabet. However, the order of the letters is unknown to you.

You are given a list of strings `words` from the alien language's dictionary, where the strings in `words` are **sorted lexicographically** by the rules of this new language.

Derive the order of letters in this language, and return it. If the given input is invalid, return `""`. If there are multiple valid solutions, return **any** of them.

## Examples

### Example 1
```
Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"
```

### Example 2
```
Input: words = ["z","x"]
Output: "zx"
```

### Example 3
```
Input: words = ["z","x","z"]
Output: ""
Explanation: The order is invalid, so return "".
```

## Constraints

- `1 <= words.length <= 100`
- `1 <= words[i].length <= 100`
- `words[i]` consists of only lowercase English letters.

## Approach Hints

1. **Build graph from adjacent words:** Compare adjacent words character by character to find ordering constraints (directed edges).
2. **Topological sort:** Apply Kahn's algorithm or DFS-based topological sort on the character graph.
3. **Edge cases:** Check for invalid inputs (e.g., ["abc", "ab"] where a longer word comes before its prefix).

## Related Problems

- 207. Course Schedule
- 210. Course Schedule II
- 953. Verifying an Alien Dictionary

## Google Follow-ups

- What if the dictionary is very large but the alphabet is small?
- How would you handle ties (characters with no ordering constraint)?
- Can you determine the minimum number of words needed to fully determine the order?
