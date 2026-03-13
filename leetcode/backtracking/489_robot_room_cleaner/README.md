# 489. Robot Room Cleaner

**Difficulty:** Hard

**Topics:** Backtracking, Interactive

## Description

You are controlling a robot that is located somewhere in a room. The room is modeled as an `m x n` binary grid where `0` represents a wall and `1` represents an empty slot.

The robot starts at an unknown location in the room that is guaranteed to be empty, and you do not have access to the grid, but you can use the provided Robot interface to control the robot.

Design an algorithm to clean the entire room using the following API:
- `Move()` — returns true if the cell in front is open and the robot moves into the cell; returns false if blocked.
- `TurnLeft()` — robot turns 90 degrees counterclockwise.
- `TurnRight()` — robot turns 90 degrees clockwise.
- `Clean()` — cleans the current cell.

## Examples

**Example 1:**
```
Input: room = [[1,1,1,1],[1,1,1,1]], row = 0, col = 0
Output: Robot cleans all cells
```

## Constraints

- `m == room.length`
- `n == room[i].length`
- `1 <= m, n <= 300`
- `room[i][j]` is either `0` or `1`.
- The robot's initial position is guaranteed to be empty.

## Approach Hints

1. **DFS with backtracking:** Use DFS from the starting position. At each cell, try all 4 directions. Use a set to track visited cells.
2. **Go back:** After exploring a direction, turn 180 degrees, move back, and turn 180 degrees again to restore original orientation.
3. **Relative coordinates:** Track position relative to start since you don't know absolute coordinates.
4. **Time:** O(n - m) where n is total cells and m is obstacles. **Space:** O(n - m) for visited set.

## Related Problems

- [79. Word Search](../079_word_search/)
- [200. Number of Islands](https://leetcode.com/problems/number-of-islands/)

## Google Follow-ups

- What if the robot has limited battery? How would you optimize the cleaning path?
- How would you handle multiple robots cleaning the same room?
- What if the robot's Move() has a probability of failure?
