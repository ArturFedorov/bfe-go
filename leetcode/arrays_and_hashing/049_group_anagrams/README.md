# 49. Group Anagrams

**Difficulty:** Medium

**Topics:** Array, Hash Table, String, Sorting

## Description

Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, using all the original letters exactly once.

## Examples

### Example 1
```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

### Example 2
```
Input: strs = [""]
Output: [[""]]
```

### Example 3
```
Input: strs = ["a"]
Output: [["a"]]
```

## Constraints

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Approach Hints

1. **Sorted String as Key:** Sort each string and use as hash map key. O(n * k log k) time where k is max string length.
2. **Character Count as Key:** Count character frequencies and use as key. O(n * k) time.
3. **Prime Number Product:** Map each character to a prime number and use the product as key. O(n * k) time but risk of overflow.

## Related Problems

- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)
- [249. Group Shifted Strings](https://leetcode.com/problems/group-shifted-strings/)
- [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/)

## What a Google Interviewer Would Ask Next

- **What if strings are very long?** Use character frequency count as the key instead of sorting — reduces from O(k log k) to O(k) per string.
- **What about Unicode characters?** Use a map instead of a fixed-size array for character counts to handle any character set.
- **How would you do this in a distributed system?** Use MapReduce — map each string to (sorted_key, string) pairs, then reduce by key to collect groups.
- **What's the time complexity difference between sorting vs counting?** Sorting: O(n * k log k). Counting: O(n * k). Counting wins when strings are long.
