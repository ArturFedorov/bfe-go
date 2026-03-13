# 253. Meeting Rooms II

**Difficulty:** Medium

**Topics:** Array, Two Pointers, Greedy, Sorting, Heap (Priority Queue)

## Description

Given an array of meeting time intervals `intervals` where `intervals[i] = [starti, endi]`, return the minimum number of conference rooms required.

## Examples

**Example 1:**
```
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2
```

**Example 2:**
```
Input: intervals = [[7,10],[2,4]]
Output: 1
```

## Constraints

- `1 <= intervals.length <= 10^4`
- `0 <= starti < endi <= 10^6`

## Approach Hints

1. **Min-heap:** Sort by start time, use a min-heap to track end times of ongoing meetings. If the earliest ending meeting ends before the current one starts, reuse that room.
2. **Chronological ordering:** Separate start and end times, sort them, use two pointers to count overlapping meetings.
3. **Time:** O(n log n). **Space:** O(n).

## Related Problems

- [56. Merge Intervals](../056_merge_intervals/)
- [57. Insert Interval](../057_insert_interval/)
- [759. Employee Free Time](../759_employee_free_time/)

## Google Follow-ups

- What if meetings have priorities and you have a limited number of rooms?
- How would you find which specific meetings go in which rooms?
- Can you handle dynamic additions/removals of meetings?
