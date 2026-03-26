# 706. Design HashMap

**Difficulty:** Easy

**Topics:** Array, Hash Table, Linked List, Design, Hash Function

## Description

Design a HashMap without using any built-in hash table libraries.

Implement the `MyHashMap` class:

- `MyHashMap()` Initializes the object with an empty map.
- `void put(int key, int value)` Inserts a `(key, value)` pair into the HashMap. If the `key` already exists in the map, update the corresponding `value`.
- `int get(int key)` Returns the `value` to which the specified `key` is mapped, or `-1` if this map contains no mapping for the `key`.
- `void remove(int key)` Removes the `key` and its corresponding `value` if the map contains the mapping for the `key`.

## Examples

**Example 1:**

```
Input: ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
       [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Output: [null, null, null, 1, -1, null, 1, null, -1]
```

## Constraints

- `0 <= key, value <= 10^6`
- At most `10^4` calls will be made to `put`, `get`, and `remove`.

## Approach Hints

1. **Fixed-size array:** Use a large array (size 10^6+1) for direct addressing. O(1) all operations but high memory.
2. **Chaining with linked lists:** Use an array of buckets with linked lists for collision resolution. Good balance of time and space.
3. **Open addressing:** Use a flat array with probing (linear, quadratic, or double hashing) to resolve collisions.

## Related Problems

- [705. Design HashSet](https://leetcode.com/problems/design-hashset/)

## Google Follow-ups

- How would you handle dynamic resizing (rehashing)?
- How would you make this thread-safe?
- What hash function would you choose and why?
