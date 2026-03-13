# 127. Word Ladder

**Difficulty:** Hard

**Topics:** Hash Table, String, Breadth-First Search

## Description

A **transformation sequence** from word `beginWord` to word `endWord` using a dictionary `wordList` is a sequence of words `beginWord -> s1 -> s2 -> ... -> sk` such that:

- Every adjacent pair of words differs by a single letter.
- Every `si` for `1 <= i <= k` is in `wordList`. Note that `beginWord` does not need to be in `wordList`.
- `sk == endWord`

Given two words, `beginWord` and `endWord`, and a dictionary `wordList`, return the **number of words** in the **shortest transformation sequence** from `beginWord` to `endWord`, or `0` if no such sequence exists.

## Examples

### Example 1
```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> "cog", which is 5 words long.
```

### Example 2
```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
```

## Constraints

- `1 <= beginWord.length <= 10`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`, `endWord`, and `wordList[i]` consist of lowercase English letters.
- `beginWord != endWord`
- All the words in `wordList` are **unique**.

## Approach Hints

1. **BFS:** Treat each word as a node. Two words are connected if they differ by one letter. BFS finds the shortest path.
2. **Wildcard patterns:** For each word, generate patterns like `h*t`, `*ot`, `ho*` and group words by pattern for O(1) neighbor lookup.
3. **Bidirectional BFS:** Start BFS from both ends simultaneously for significant speedup.

## Related Problems

- 126. Word Ladder II
- 433. Minimum Genetic Mutation
- 752. Open the Lock

## Google Follow-ups

- How would you return all shortest transformation sequences?
- Can you optimize for very large dictionaries with millions of words?
- What if words can have different lengths (insertions/deletions allowed)?
