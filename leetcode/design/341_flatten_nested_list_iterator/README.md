# 341. Flatten Nested List Iterator

**Difficulty:** Medium

**Topics:** Stack, Tree, Depth-First Search, Design, Queue, Iterator

## Description

You are given a nested list of integers `nestedList`. Each element is either an integer or a list whose elements may also be integers or other lists. Implement an iterator to flatten it.

Implement the `NestedIterator` class:

- `NestedIterator(List<NestedInteger> nestedList)` Initializes the iterator with the nested list `nestedList`.
- `int next()` Returns the next integer in the nested list.
- `boolean hasNext()` Returns `true` if there are still some integers in the nested list.

## Examples

**Example 1:**

```
Input: nestedList = [[1,1],2,[1,1]]
Output: [1,1,2,1,1]
```

**Example 2:**

```
Input: nestedList = [1,[4,[6]]]
Output: [1,4,6]
```

## Constraints

- `1 <= nestedList.length <= 500`
- The values of the integers in the nested list are in the range `[-10^6, 10^6]`.

## Approach Hints

1. **Flatten eagerly in constructor:** DFS to collect all integers into a slice, then iterate over it.
2. **Lazy flattening with stack:** Push elements onto a stack in reverse order; in `HasNext`, keep flattening the top until it is an integer.
3. **Recursive generator style:** Use a queue/channel approach for lazy evaluation.

## Related Problems

- [251. Flatten 2D Vector](https://leetcode.com/problems/flatten-2d-vector/)
- [281. Zigzag Iterator](https://leetcode.com/problems/zigzag-iterator/)
- [385. Mini Parser](https://leetcode.com/problems/mini-parser/)

## Google Follow-ups

- What if the nesting depth can be very large (millions of levels)?
- How would you implement this as a truly lazy iterator without pre-flattening?
- What if the nested structure can contain cycles?
