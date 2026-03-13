# 22. Generate Parentheses

**Difficulty:** Medium

**Topics:** String, Dynamic Programming, Backtracking

## Description

Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

## Examples

**Example 1:**
```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

**Example 2:**
```
Input: n = 1
Output: ["()"]
```

**Example 3:**
```
Input: n = 0
Output: []
```

## Constraints

- `0 <= n <= 8`

## Approach Hints

1. **Backtracking:** Track count of open and close parens used. Add `(` if open < n, add `)` if close < open.
2. **DP/Catalan:** Use the recurrence relation based on Catalan numbers to build results for `n` from smaller values.
3. **BFS:** Generate level by level, pruning invalid states.

## Related Problems

- [17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)
- [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
- [301. Remove Invalid Parentheses](https://leetcode.com/problems/remove-invalid-parentheses/)

## Google Follow-ups

- How would you generate only the k-th valid combination without generating all?
- Extend to multiple types of brackets: `()`, `[]`, `{}`.
- What is the time complexity in terms of the Catalan number?
