# 621. Task Scheduler

**Difficulty:** Medium

**Topics:** Array, Hash Table, Greedy, Sorting, Heap (Priority Queue), Counting

## Description

You are given an array of CPU `tasks`, each represented by letters A to Z, and a cooling interval `n`. Each cycle or interval allows the completion of one task. Tasks can be completed in any order, but there's a constraint: identical tasks must be separated by at least `n` intervals due to cooling time.

Return the minimum number of intervals the CPU will take to finish all the given tasks.

## Examples

**Example 1:**
```
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: A->B->idle->A->B->idle->A->B
```

**Example 2:**
```
Input: tasks = ["A","A","A","B","B","B"], n = 0
Output: 6
Explanation: No cooling needed, just run all tasks.
```

**Example 3:**
```
Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
Output: 16
Explanation: A->B->C->A->D->E->A->F->G->A->idle->idle->A->idle->idle->A
```

## Constraints

- `1 <= tasks.length <= 10^4`
- `tasks[i]` is an uppercase English letter.
- `0 <= n <= 100`

## Approach Hints

1. **Greedy formula:** Count the max frequency `f`. The minimum time is `max(len(tasks), (f-1)*(n+1) + count_of_tasks_with_max_freq)`.
2. **Heap simulation:** Use a max-heap to always schedule the most frequent remaining task, with a cooldown queue.
3. **Time:** O(n) for formula approach. **Space:** O(1) since at most 26 different tasks.

## Related Problems

- [347. Top K Frequent Elements](../347_top_k_frequent/)
- [358. Rearrange String k Distance Apart](https://leetcode.com/problems/rearrange-string-k-distance-apart/)

## Google Follow-ups

- What if tasks have different execution times?
- What if there are task dependencies (DAG)?
- How would you actually output the optimal schedule, not just the length?
