# 904. Fruit Into Baskets

**Difficulty:** Medium

**Topics:** Array, Hash Table, Sliding Window

## Description

You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array `fruits` where `fruits[i]` is the type of fruit the `i`-th tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules:

- You only have two baskets, and each basket can only hold a single type of fruit.
- There is no limit on the amount of fruit each basket can hold.
- Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. You must stop when you encounter a tree with fruit that cannot fit in either basket.

Return the maximum number of fruits you can pick.

## Examples

### Example 1
```
Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.
```

### Example 2
```
Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2]. If we started at tree 0, we would only pick [0,1].
```

### Example 3
```
Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2]. If we started at tree 0, we would only pick [1,2].
```

## Constraints

- `1 <= fruits.length <= 10^5`
- `0 <= fruits[i] < fruits.length`

## Approach Hints

1. This is essentially "longest subarray with at most 2 distinct elements" — a classic sliding window problem.
2. Maintain a hash map counting fruit types in the current window. Expand the right pointer; when distinct types exceed 2, shrink from the left.
3. Track the maximum window size throughout the process.

## Related Problems

- [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
- [159. Longest Substring with At Most Two Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/)
- [340. Longest Substring with At Most K Distinct Characters](https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/)

## What a Google Interviewer Would Ask Next

- **How would you generalize this to k baskets?** Replace the hardcoded 2 with k; the sliding window logic is identical, just check map size against k.
- **What's the time complexity?** O(n) since each element is added and removed from the window at most once.
- **Could you solve this without a hash map?** For exactly 2 types you can track the two types and the last contiguous run, but the hash map approach is cleaner and generalizes.
