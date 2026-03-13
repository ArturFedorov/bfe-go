# 20. Valid Parentheses

**Difficulty:** Easy

**Topics:** String, Stack

## Description

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Examples

### Example 1
```
Input: s = "()"
Output: true
```

### Example 2
```
Input: s = "()[]{}"
Output: true
```

### Example 3
```
Input: s = "(]"
Output: false
```

### Example 4
```
Input: s = "([)]"
Output: false
```

### Example 5
```
Input: s = "{[]}"
Output: true
```

## Constraints

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`.

## Approach Hints

1. Use a stack: push opening brackets, pop on closing brackets and check for matching.
2. If the stack is empty when encountering a closing bracket, or the top doesn't match, return false.
3. After processing all characters, the string is valid only if the stack is empty.

## Related Problems

- [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)
- [32. Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)
- [1249. Minimum Remove to Make Valid Parentheses](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/)

## What a Google Interviewer Would Ask Next

- **What if the string also contains non-bracket characters?** Simply skip characters that aren't brackets.
- **How would you find the longest valid parentheses substring?** Use a stack storing indices, or DP where dp[i] = length of longest valid substring ending at i.
- **What if you had custom bracket types defined at runtime?** Use a map from closing to opening bracket, making the solution generic.
