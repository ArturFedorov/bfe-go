# 681. Next Closest Time

**Difficulty:** Medium

**Topics:** String, Enumeration, Simulation

## Description

Given a time represented in the format `"HH:MM"`, form the next closest time by reusing the current digits. There is no limit on how many times a digit can be reused.

You may assume the given input string is always valid. For example, `"01:34"`, `"12:09"` are all valid. `"1:34"`, `"12:9"` are all invalid.

## Examples

### Example 1
```
Input: time = "19:34"
Output: "19:39"
Explanation: The next closest time choosing from digits [1, 9, 3, 4] is 19:39, which occurs 5 minutes later.
```

### Example 2
```
Input: time = "23:59"
Output: "22:22"
Explanation: The next closest time choosing from digits [2, 3, 5, 9] is 22:22. It may be assumed that the returned time is next day's time since it is smaller than the input time numerically.
```

## Constraints

- Input is a valid time in `"HH:MM"` format.
- `HH` is in range `[00, 23]`.
- `MM` is in range `[00, 59]`.

## Approach Hints

1. Brute force: simulate minute by minute (at most 1440 steps) and check if each time uses only the original digits.
2. Enumerate all valid times from the 4 digits (at most 4^4 = 256 combinations), filter valid ones, and find the next time after the current.
3. Greedy digit replacement from right to left: try to increment each position with the smallest valid digit that is larger than the current one.

## Related Problems

- [949. Largest Time for Given Digits](https://leetcode.com/problems/largest-time-for-given-digits/)
- [539. Minimum Time Difference](https://leetcode.com/problems/minimum-time-difference/)

## What a Google Interviewer Would Ask Next

- **What's the time complexity of the simulation approach?** O(1) since we iterate at most 1440 minutes, and checking digits is constant. Everything is bounded by constants.
- **How would you handle 12-hour format with AM/PM?** Add AM/PM logic and wrap at 12:59 AM/PM boundaries; the digit-reuse constraint stays the same.
- **What if you could also reuse digits from a given pool, not just the current time?** Same enumeration approach; just change the digit set to the given pool.
