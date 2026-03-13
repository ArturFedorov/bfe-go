# 409. Longest Palindrome

**Difficulty:** Easy

**Topics:** Hash Table, String, Greedy

## Description

Given a string `s` which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, `"Aa"` is not considered a palindrome.

## Examples

### Example 1
```
Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
```

### Example 2
```
Input: s = "a"
Output: 1
```

## Constraints

- `1 <= s.length <= 2000`
- `s` consists of lowercase and/or uppercase English letters only.

## Approach Hints

1. Count the frequency of each character. Every character with an even count contributes fully. For odd counts, contribute count - 1.
2. If any character has an odd count, you can place one in the center, so add 1 to the total.
3. This is a greedy approach with O(n) time and O(1) space (fixed alphabet size).

## Related Problems

- [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)
- [266. Palindrome Permutation](https://leetcode.com/problems/palindrome-permutation/)
- [516. Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/)

## What a Google Interviewer Would Ask Next

- **What if you need to construct the actual palindrome, not just the length?** Greedily place half the even-count characters, add one odd character in the center if available, then mirror.
- **What if the input is a stream of characters?** Maintain a running frequency map and a count of odd-frequency characters to update the answer in O(1) per new character.
- **How does this change if the string contains Unicode?** The algorithm is identical; just use a map[rune]int instead of a fixed array.
