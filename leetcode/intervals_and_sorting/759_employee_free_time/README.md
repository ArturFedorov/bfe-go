# 759. Employee Free Time

**Difficulty:** Hard

**Topics:** Array, Sorting, Heap (Priority Queue)

## Description

We are given a list `schedule` of employees, which represents the working time for each employee. Each employee has a list of non-overlapping `Intervals`, and these intervals are in sorted order.

Return the list of finite intervals representing common, positive-length free time for all employees, also in sorted order.

## Examples

**Example 1:**
```
Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
Output: [[3,4]]
Explanation: There are a total of three employees, and all common
free time intervals would be [-inf,1],[3,4],[10,inf].
We discard any intervals that contain inf as they aren't finite.
```

**Example 2:**
```
Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
Output: [[5,6],[7,9]]
```

## Constraints

- `1 <= schedule.length, schedule[i].length <= 50`
- `0 <= schedule[i][j].Start < schedule[i][j].End <= 10^8`

## Approach Hints

1. **Flatten and merge:** Collect all intervals from all employees, merge them (like problem 56), then find gaps between merged intervals.
2. **Min-heap:** Use a heap to process intervals in order across all employees.
3. **Time:** O(n log n) where n is total number of intervals. **Space:** O(n).

## Related Problems

- [56. Merge Intervals](../056_merge_intervals/)
- [253. Meeting Rooms II](../253_meeting_rooms_ii/)

## Google Follow-ups

- What if employees are in different time zones?
- How would you find the longest common free time window of at least K duration?
- Can you solve this in a streaming fashion as employee schedules update?
