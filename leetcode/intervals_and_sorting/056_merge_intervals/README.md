# 56. Merge Intervals

**Difficulty:** Medium

**Topics:** Array, Sorting

## Description

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

## Examples

**Example 1:**
```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
```

**Example 2:**
```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

## Constraints

- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^4`

## Approach Hints

1. **Sort first:** Sort intervals by start time. Overlapping intervals will be adjacent after sorting.
2. **Merge greedily:** Iterate through sorted intervals, extending the current interval's end or starting a new one.
3. **Time:** O(n log n) for sorting. **Space:** O(n) for result.

## Related Problems

- [57. Insert Interval](../057_insert_interval/)
- [253. Meeting Rooms II](../253_meeting_rooms_ii/)
- [759. Employee Free Time](../759_employee_free_time/)

## Google Follow-ups

- What if intervals are streamed one at a time? How would you maintain merged intervals online?
- How would you handle this with very large datasets that don't fit in memory?
- Can you do this in-place without extra space for the result?
