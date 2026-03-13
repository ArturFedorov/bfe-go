# 44. Wildcard Matching

**Difficulty:** Hard

**Topics:** String, Dynamic Programming, Greedy, Recursion

## Description

Given an input string `s` and a pattern `p`, implement wildcard pattern matching with support for `'?'` and `'*'` where:

- `'?'` Matches any single character.
- `'*'` Matches any sequence of characters (including the empty sequence).

The matching should cover the **entire** input string (not partial).

## Examples

**Example 1:**
```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

**Example 2:**
```
Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.
```

**Example 3:**
```
Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'b', not 'a'.
```

**Example 4:**
```
Input: s = "adceb", p = "*a*b"
Output: true
Explanation: '*' matches "adce" and 'b' matches 'b'.
```

## Constraints

- `0 <= s.length, p.length <= 2000`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'?'` or `'*'`.

## Approach Hints

1. **DP:** Build `dp[i][j]` = whether `s[0..i)` matches `p[0..j)`. Handle `*` by combining "match empty" and "match one more char".
2. **Greedy with backtracking:** Track the last `*` position and backtrack when a mismatch occurs.
3. **Two-pointer:** Linear time approach using star/match index tracking.

## Related Problems

- [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)
- [1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)

## Google Follow-ups

- Can you solve it in O(n) space using a rolling array?
- How does this differ from regex matching, and which is harder?
- How would you handle `**` meaning "match across path separators" (like glob)?
