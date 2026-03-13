# 139. Word Break

**Difficulty:** Medium

**Topics:** Array, Hash Table, String, Dynamic Programming, Trie, Memoization

## Description

Given a string `s` and a dictionary of strings `wordDict`, return `true` if `s` can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

## Examples

**Example 1:**
```
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
```

**Example 2:**
```
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
```

**Example 3:**
```
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```

## Constraints

- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` and `wordDict[i]` consist of only lowercase English letters.
- All the strings of `wordDict` are **unique**.

## Approach Hints

1. **DP:** `dp[i]` = whether `s[0..i)` can be segmented. For each `i`, check all words that could end at position `i`.
2. **BFS/DFS:** Treat each valid prefix as a node, explore further segmentations.
3. **Trie:** Build a trie from `wordDict` for efficient prefix matching during DP.

## Related Problems

- [140. Word Break II](https://leetcode.com/problems/word-break-ii/)
- [472. Concatenated Words](https://leetcode.com/problems/concatenated-words/)

## Google Follow-ups

- Return all possible segmentations (Word Break II).
- What if the dictionary is very large — how would you optimize?
- How would you handle this with a trie for prefix matching?
