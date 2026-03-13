# 642. Design Search Autocomplete System

**Difficulty:** Hard

**Topics:** String, Design, Trie, Sorting, Heap (Priority Queue), Data Stream

## Description

Design a search autocomplete system for a search engine. Users may input a sentence (at least one word and end with a special character `'#'`).

You are given a string array `sentences` and an integer array `times` both of length `n` where `sentences[i]` is a previously typed sentence and `times[i]` is the corresponding number of times the sentence was typed. For each input character except `'#'`, return the top 3 historical hot sentences that have the same prefix as the part of the sentence already typed.

Here are the specific rules:

- The hot degree for a sentence is defined as the number of times a user typed the exactly same sentence before.
- The returned top 3 hot sentences should be sorted by hot degree (descending). If several sentences have the same hot degree, use ASCII-code order (ascending).
- If less than 3 hot sentences exist, return as many as you can.
- When the input is `'#'`, it means the sentence ends, and in this case, you need to return an empty list and store the input sentence.

## Examples

**Example 1:**

```
Input: ["AutocompleteSystem", "input", "input", "input", "input"]
       [[["i love you", "island", "iroman", "i love leetcode"], [5, 3, 2, 2]], ['i'], [' '], ['a'], ['#']]
Output: [null, ["i love you", "island", "i love leetcode"], ["i love you", "i love leetcode"], [], []]
```

## Constraints

- `n == sentences.length`
- `n == times.length`
- `1 <= n <= 100`
- `1 <= sentences[i].length <= 100`
- `1 <= times[i] <= 50`
- `sentences[i]` consists of lowercase English letters and spaces.
- The input character will be a lowercase letter or `'#'`.
- At most `5000` calls will be made to `input`.

## Approach Hints

1. **Trie + sorting:** Store sentences in a trie. At each node, keep a map of sentence to count. On input, traverse the trie and sort candidates.
2. **Trie + heap:** Same as above but use a min-heap of size 3 to efficiently get top 3.
3. **Hash map with prefix:** For each sentence, store all its prefixes mapped to the sentence. Simple but uses more memory.

## Related Problems

- [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/)
- [211. Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/)

## Google Follow-ups

- How would you handle real-time ranking updates across distributed servers?
- What if you want to support typo correction / fuzzy matching?
- How would you personalize autocomplete results per user?
- How would you handle offensive or sensitive content filtering?
