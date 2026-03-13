# 359. Logger Rate Limiter

**Difficulty:** Easy

**Topics:** Hash Table, Design

## Description

Design a logger system that receives a stream of messages along with their timestamps. Each unique message should only be printed at most every 10 seconds (i.e. a message printed at timestamp `t` will prevent the same message from being printed until timestamp `t + 10`).

All messages will come in chronological order. Several messages may arrive at the same timestamp.

Implement the `Logger` class:

- `Logger()` Initializes the logger object.
- `bool shouldPrintMessage(int timestamp, string message)` Returns `true` if the message should be printed in the given timestamp, otherwise returns `false`.

## Examples

**Example 1:**

```
Input: ["Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"]
       [[], [1, "foo"], [2, "bar"], [3, "foo"], [8, "bar"], [10, "foo"], [11, "foo"]]
Output: [null, true, true, false, false, false, true]
```

## Constraints

- `0 <= timestamp <= 10^9`
- Every timestamp will be passed in non-decreasing order (chronological order).
- `1 <= message.length <= 30`
- At most `10^4` calls will be made to `shouldPrintMessage`.

## Approach Hints

1. **Hash map:** Store the last allowed timestamp for each message. O(1) per call.
2. **Sliding window with queue + set:** Keep a queue of `(timestamp, message)` pairs and a set of active messages. Clean up expired entries.
3. **Circular buffer:** Use a fixed-size array of 10 buckets (one per second mod 10), each with a set of messages.

## Related Problems

- [353. Design Snake Game](https://leetcode.com/problems/design-snake-game/)
- [362. Design Hit Counter](https://leetcode.com/problems/design-hit-counter/)

## Google Follow-ups

- How would you handle this in a multi-threaded environment?
- What if memory is limited and you have millions of unique messages?
- How would you modify this if timestamps are not guaranteed to be in order?
