# 155. Min Stack

**Difficulty:** Medium

**Topics:** Stack, Design

## Description

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

You must implement a solution with `O(1)` time complexity for each function.

## Examples

**Example 1:**

```
Input: ["MinStack","push","push","push","getMin","pop","top","getMin"]
       [[],[-2],[0],[-3],[],[],[],[]]
Output: [null,null,null,null,-3,null,0,-2]
```

## Constraints

- `-2^31 <= val <= 2^31 - 1`
- Methods `pop`, `top` and `getMin` operations will always be called on non-empty stacks.
- At most `3 * 10^4` calls will be made to `push`, `pop`, `top`, and `getMin`.

## Approach Hints

1. **Two stacks:** Maintain a second stack that tracks the current minimum at each level.
2. **Single stack with pairs:** Store `(value, currentMin)` pairs in one stack.
3. **Encoding trick:** Store the difference between the value and the current min to avoid extra space.

## Related Problems

- [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)
- [716. Max Stack](https://leetcode.com/problems/max-stack/)

## Google Follow-ups

- How would you handle `getMin` in O(1) time and O(1) extra space?
- What if you need to support `getMax` as well?
- How would you design a thread-safe version?
