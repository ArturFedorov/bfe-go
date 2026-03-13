# 438. Find All Anagrams in a String

**Difficulty:** Medium

**Topics:** Hash Table, String, Sliding Window

## Description

Given two strings `s` and `p`, return an array of all the start indices of `p`'s anagrams in `s`. You may return the answer in any order.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Examples

### Example 1
```
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
```

### Example 2
```
Input: s = "abab", p = "ab"
Output: [0,1,2]
```

## Constraints

- `1 <= s.length, p.length <= 3 * 10^4`
- `s` and `p` consist of lowercase English letters.

## Approach Hints

1. Use a fixed-size sliding window of length len(p) over s, maintaining character frequency counts.
2. Compare the window's frequency array with p's frequency array; if they match, record the start index.
3. Optimize by tracking how many characters have matching counts (a "matches" counter) to avoid comparing all 26 entries each step.

## Related Problems

- [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)
- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)
- [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/)

## What a Google Interviewer Would Ask Next

- **How is this different from problem 567 (Permutation in String)?** It's essentially the same problem, but 567 asks for a boolean while 438 asks for all starting indices.
- **Can you solve this in O(n) with O(1) space?** Yes, since the alphabet is fixed at 26 lowercase letters, the frequency arrays are constant size.
- **What if the characters were Unicode instead of just lowercase letters?** Replace the [26]int array with a map[rune]int; the sliding window logic stays the same.
