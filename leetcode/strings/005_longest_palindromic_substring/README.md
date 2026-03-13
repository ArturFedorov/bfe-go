# 5. Longest Palindromic Substring

**Difficulty:** Medium

**Topics:** String, Dynamic Programming, Two Pointers

## Description

Given a string `s`, return the longest palindromic substring in `s`.

## Examples

### Example 1
```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

### Example 2
```
Input: s = "cbbd"
Output: "bb"
```

## Constraints

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## Approach Hints

1. Expand around center: for each index (and each pair of adjacent indices), expand outward while characters match. Track the longest palindrome found.
2. Dynamic programming: dp[i][j] = true if s[i..j] is a palindrome. Build from shorter substrings to longer ones.
3. Manacher's algorithm achieves O(n) time by reusing previously computed palindrome information, but expand-around-center O(n^2) is usually sufficient for interviews.

## Related Problems

- [516. Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/)
- [647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)
- [214. Shortest Palindrome](https://leetcode.com/problems/shortest-palindrome/)

## What a Google Interviewer Would Ask Next

- **Can you explain Manacher's algorithm?** It processes the string linearly, using a center and right boundary to skip redundant comparisons by leveraging symmetry of already-found palindromes.
- **What if you need to find all maximal palindromic substrings?** Manacher's directly gives this; expand-around-center can also enumerate them in O(n^2).
- **How would you handle this for very long strings (10^6+)?** Manacher's O(n) or suffix array/tree approaches become necessary.
