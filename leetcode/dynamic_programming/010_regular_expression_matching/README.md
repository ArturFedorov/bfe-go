# 10. Regular Expression Matching

**Difficulty:** Hard

**Topics:** String, Dynamic Programming, Recursion

## Description

Given an input string `s` and a pattern `p`, implement regular expression matching with support for `'.'` and `'*'` where:

- `'.'` Matches any single character.
- `'*'` Matches zero or more of the preceding element.

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
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

**Example 3:**
```
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

**Example 4:**
```
Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c is repeated 0 times, a is repeated twice, followed by b.
```

## Constraints

- `1 <= s.length <= 20`
- `1 <= p.length <= 20`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'.'`, and `'*'`.
- It is guaranteed for each appearance of the character `'*'`, there will be a previous valid character to match.

## Approach Hints

1. **Recursive:** Process pattern from left to right, handling `x*` pairs specially (match zero or more).
2. **Top-down DP:** Memoize recursive calls with `(i, j)` indices into `s` and `p`.
3. **Bottom-up DP:** Build a 2D table `dp[i][j]` = whether `s[0..i)` matches `p[0..j)`.

## Related Problems

- [44. Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)
- [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

## Google Follow-ups

- How would you extend this to support `+` (one or more) and `?` (zero or one)?
- Can you implement this with an NFA (nondeterministic finite automaton)?
- What if the pattern could include character classes like `[a-z]`?
