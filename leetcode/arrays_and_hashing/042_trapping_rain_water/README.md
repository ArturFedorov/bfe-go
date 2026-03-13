# 42. Trapping Rain Water

**Difficulty:** Hard

**Topics:** Array, Two Pointers, Dynamic Programming, Stack, Monotonic Stack

## Description

Given `n` non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

## Examples

### Example 1
```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water are being trapped.
```

### Example 2
```
Input: height = [4,2,0,3,2,5]
Output: 9
```

## Constraints

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## Approach Hints

1. Two-pointer approach: maintain left and right pointers with running maxLeft and maxRight. Water at each position is determined by the smaller of the two maxes minus the current height.
2. DP approach: precompute prefix-max from left and suffix-max from right, then water at index i = min(leftMax[i], rightMax[i]) - height[i].
3. Monotonic stack approach: maintain a decreasing stack of indices; when a taller bar is found, pop and calculate trapped water between the current bar and the new stack top.

## Related Problems

- [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)
- [407. Trapping Rain Water II](https://leetcode.com/problems/trapping-rain-water-ii/)
- [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)

## What a Google Interviewer Would Ask Next

- **Can you extend this to 2D (Trapping Rain Water II)?** Use a min-heap BFS from the boundary inward, processing cells in order of height.
- **Which approach would you choose for a streaming input?** The stack-based approach works well since it processes elements left to right in a single pass.
- **What's the space complexity trade-off between the three approaches?** Two pointers: O(1); stack: O(n) worst case; DP: O(n) for the prefix arrays.
