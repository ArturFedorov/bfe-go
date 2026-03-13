# 362. Design Hit Counter

**Difficulty:** Medium

**Topics:** Array, Binary Search, Design, Queue

## Description

Design a hit counter which counts the number of hits received in the past 5 minutes (i.e., the past `300` seconds).

Your system should accept a `timestamp` parameter (in seconds granularity), and you may assume that calls are being made to the system in chronological order (i.e., `timestamp` is monotonically increasing). Several hits may arrive at roughly the same time.

Implement the `HitCounter` class:

- `HitCounter()` Initializes the object of the hit counter system.
- `void hit(int timestamp)` Records a hit that happened at `timestamp` (in seconds). Several hits may happen at the same `timestamp`.
- `int getHits(int timestamp)` Returns the number of hits in the past 5 minutes from `timestamp` (i.e., the past `300` seconds).

## Examples

**Example 1:**

```
Input: ["HitCounter", "hit", "hit", "hit", "getHits", "hit", "getHits", "getHits"]
       [[], [1], [2], [3], [4], [300], [300], [301]]
Output: [null, null, null, null, 3, null, 4, 3]
```

## Constraints

- `1 <= timestamp <= 2 * 10^9`
- All the calls are being made to the system in chronological order (i.e., `timestamp` is monotonically increasing).
- At most `300` calls will be made to `hit` and `getHits`.

## Approach Hints

1. **Queue:** Store all timestamps in a queue; dequeue entries older than 300 seconds on each `getHits` call.
2. **Circular buffer:** Use arrays of size 300 for timestamps and counts. O(1) `hit`, O(300) `getHits`.
3. **Binary search:** Store sorted timestamps and use binary search to find the count in the window.

## Related Problems

- [359. Logger Rate Limiter](https://leetcode.com/problems/logger-rate-limiter/)

## Google Follow-ups

- How would you handle a very high volume of hits per second?
- How would you scale this across multiple machines?
- What if timestamps can arrive out of order?
