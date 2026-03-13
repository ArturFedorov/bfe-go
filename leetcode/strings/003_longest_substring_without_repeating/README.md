# 3. Longest Substring Without Repeating Characters

**Difficulty:** Medium

**Topics:** Hash Table, String, Sliding Window

## Description

Given a string `s`, find the length of the longest substring without repeating characters.

## Examples

### Example 1
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

### Example 2
```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

### Example 3
```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
```

## Constraints

- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols and spaces.

## Approach Hints

1. Use a sliding window with a hash map storing the last index of each character.
2. When a duplicate is found, jump the left pointer to max(left, lastIndex[char] + 1) to skip past the previous occurrence.
3. Update the maximum length at each step as right - left + 1.

## Related Problems

- [159. Longest Substring with At Most Two Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/)
- [340. Longest Substring with At Most K Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/)
- [992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/)

## What a Google Interviewer Would Ask Next

- **What if we need the actual substring, not just the length?** Track the start index of the best window and return s[bestStart:bestStart+maxLen].
- **How does this handle Unicode/multibyte characters?** In Go, iterating with range gives runes; using a map[rune]int handles Unicode correctly.
- **Can you do this with a fixed-size array instead of a hash map?** Yes, if the charset is limited (e.g., ASCII), use a [128]int array for O(1) lookups with better constant factors.
