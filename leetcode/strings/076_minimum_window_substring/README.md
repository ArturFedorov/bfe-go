# 76. Minimum Window Substring

**Difficulty:** Hard

**Topics:** Hash Table, String, Sliding Window

## Description

Given two strings `s` and `t` of lengths `m` and `n` respectively, return the minimum window substring of `s` such that every character in `t` (including duplicates) is included in the window. If there is no such substring, return the empty string `""`.

The testcases will be generated such that the answer is unique.

## Examples

### Example 1
```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
```

### Example 2
```
Input: s = "a", t = "a"
Output: "a"
```

### Example 3
```
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window. Since s only has one 'a', return "".
```

## Constraints

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= 10^5`
- `s` and `t` consist of uppercase and lowercase English letters.

## Approach Hints

1. Build a frequency map for `t`. Use a sliding window over `s`, expanding right to include characters and shrinking left when all characters are satisfied.
2. Maintain a counter of how many distinct characters have met their required frequency to know when the window is valid.
3. When the window is valid, try to shrink from the left to find the minimum window, updating the best result.

## Related Problems

- [30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)
- [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)
- [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)

## What a Google Interviewer Would Ask Next

- **Can you do this in O(m + n)?** Yes, the sliding window approach visits each character at most twice (once by right, once by left pointer).
- **What if t contains characters not in s?** The algorithm naturally returns "" since the required count can never be satisfied.
- **How would you find all minimum windows (not just one)?** Track all windows of minimum length during the shrink phase instead of just the first one found.
