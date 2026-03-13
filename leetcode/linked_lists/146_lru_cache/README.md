# 146. LRU Cache

**Difficulty:** Medium

**Topics:** Hash Table, Linked List, Design, Doubly-Linked List

---

## Description

Design a data structure that follows the constraints of a **[Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)**.

Implement the `LRUCache` struct:

- `Constructor(capacity int) LRUCache` — Initialize the LRU cache with **positive** size `capacity`.
- `Get(key int) int` — Return the value of the `key` if the key exists, otherwise return `-1`.
- `Put(key int, value int)` — Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, **evict** the least recently used key.

The functions `Get` and `Put` must each run in **O(1)** average time complexity.

---

## Examples

### Example 1

```
Input
["Constructor", "Put", "Put", "Get", "Put", "Get", "Put", "Get", "Get", "Get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]

Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
cache := Constructor(2)
cache.Put(1, 1)   // cache is {1=1}
cache.Put(2, 2)   // cache is {1=1, 2=2}
cache.Get(1)      // return 1
cache.Put(3, 3)   // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
cache.Get(2)      // returns -1 (not found)
cache.Put(4, 4)   // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
cache.Get(1)      // return -1 (not found)
cache.Get(3)      // return 3
cache.Get(4)      // return 4
```

---

## Constraints

- `1 <= capacity <= 3000`
- `0 <= key <= 10^4`
- `0 <= value <= 10^5`
- At most `2 * 10^5` calls will be made to `Get` and `Put`.

---

## Follow-up

Could you implement the LRU cache without using the built-in `container/list` package?

---

## Approach Hints

<details>
<summary>Hint 1</summary>
What data structure gives you O(1) lookup by key?
</details>

<details>
<summary>Hint 2</summary>
What data structure lets you efficiently track insertion/access order and remove/add elements in O(1)?
</details>

<details>
<summary>Hint 3</summary>
Combine a hash map with a doubly linked list. The map stores key -> node pointers, and the linked list maintains access order.
</details>

---

## Related Problems

- [460. LFU Cache](https://leetcode.com/problems/lfu-cache/) (Hard)
- [588. Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) (Hard)
- [1472. Design Browser History](https://leetcode.com/problems/design-browser-history/) (Medium)

### What a Google Interviewer Would Ask Next
```
1. "How would you make this thread safe?"
   → sync.RWMutex around Get (RLock) and Put (Lock)
   → or sharded LRU for higher throughput

2. "How would you add TTL (time to expiry)?"
   → add expiredAt time.Time to entry struct
   → check on Get, lazy eviction

3. "How would you scale this to distributed cache?"
   → consistent hashing across LRU cache nodes
   → combines Q13 + Q14 ✅

4. "What's the memory overhead per entry?"
   → entry struct: 16 bytes
   → list.Element: 32 bytes (value + two pointers)
   → map entry: ~50 bytes
   → total: ~100 bytes per cached item