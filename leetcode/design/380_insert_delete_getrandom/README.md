# 380. Insert Delete GetRandom O(1)

**Difficulty:** Medium

**Topics:** Array, Hash Table, Math, Design, Randomized

## Description

Implement the `RandomizedSet` class:

- `RandomizedSet()` Initializes the `RandomizedSet` object.
- `bool insert(int val)` Inserts an item `val` into the set if not present. Returns `true` if the item was not present, `false` otherwise.
- `bool remove(int val)` Removes an item `val` from the set if present. Returns `true` if the item was present, `false` otherwise.
- `int getRandom()` Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average `O(1)` time complexity.

## Examples

**Example 1:**

```
Input: ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
       [[], [1], [2], [2], [], [1], [2], []]
Output: [null, true, false, true, 2, true, false, 2]
```

## Constraints

- `-2^31 <= val <= 2^31 - 1`
- At most `2 * 10^5` calls will be made to `insert`, `remove`, and `getRandom`.
- There will be at least one element in the data structure when `getRandom` is called.

## Approach Hints

1. **Hash map + dynamic array:** Map stores value-to-index. Array stores values. Swap-with-last trick for O(1) removal.
2. **Two hash maps:** One for value-to-index and one for index-to-value (less common, slightly more space).

## Related Problems

- [381. Insert Delete GetRandom O(1) - Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/)
- [710. Random Pick with Blacklist](https://leetcode.com/problems/random-pick-with-blacklist/)

## Google Follow-ups

- How would you extend this to support duplicates? (LeetCode 381)
- How would you make this thread-safe?
- What if you also need `getMin` and `getMax` in O(1)?
