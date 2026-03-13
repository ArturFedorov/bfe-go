# 17. Letter Combinations of a Phone Number

**Difficulty:** Medium

**Topics:** Hash Table, String, Backtracking

## Description

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons):
- 2: abc
- 3: def
- 4: ghi
- 5: jkl
- 6: mno
- 7: pqrs
- 8: tuv
- 9: wxyz

## Examples

### Example 1
```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

### Example 2
```
Input: digits = ""
Output: []
```

### Example 3
```
Input: digits = "2"
Output: ["a","b","c"]
```

## Constraints

- `0 <= digits.length <= 4`
- `digits[i]` is a digit in the range `['2', '9']`.

## Approach Hints

1. Use backtracking: for each digit, iterate over its mapped letters and recurse on the remaining digits.
2. Iterative BFS approach: start with [""], and for each digit, expand every existing string by appending each mapped letter.
3. The total number of combinations is the product of letter counts per digit (at most 4^4 = 256).

## Related Problems

- [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)
- [39. Combination Sum](https://leetcode.com/problems/combination-sum/)
- [46. Permutations](https://leetcode.com/problems/permutations/)

## What a Google Interviewer Would Ask Next

- **What's the time complexity?** O(4^n * n) where n is the number of digits, since each combination has length n and there are at most 4^n combinations.
- **How would you handle this if digits could be very long (e.g., 20 digits)?** You'd need to generate lazily or use an iterator pattern since 4^20 is ~10^12 combinations.
- **Can you return results in lexicographic order?** Yes, if you process letters in sorted order (which the standard mapping already is), the backtracking naturally produces sorted output.
