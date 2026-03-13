# 560. Subarray Sum Equals K

**Difficulty:** Medium

**Topics:** Array, Hash Table, Prefix Sum

---

## Description

Given an array of integers `nums` and an integer `k`, return _the total number of subarrays whose sum equals to_ `k`.

A subarray is a contiguous **non-empty** sequence of elements within an array.

---

## Examples

### Example 1

```
Input: nums = [1,1,1], k = 2
Output: 2
```

### Example 2

```
Input: nums = [1,2,3], k = 3
Output: 2
```

---

## Constraints

- `1 <= nums.length <= 2 * 10^4`
- `-1000 <= nums[i] <= 1000`
- `-10^7 <= k <= 10^7`

---

## Approach Hints

<details>
<summary>Hint 1</summary>
A brute force approach would check all subarrays, but that's O(n^2). Can you do better?
</details>

<details>
<summary>Hint 2</summary>
Think about prefix sums. If prefix[j] - prefix[i] == k, then the subarray from i+1 to j sums to k.
</details>

<details>
<summary>Hint 3</summary>
Use a hash map to store the frequency of prefix sums seen so far. For each new prefix sum, check how many times (prefix_sum - k) has appeared.
</details>

---

## Related Problems

- [1. Two Sum](https://leetcode.com/problems/two-sum/) (Easy)
- [523. Continuous Subarray Sum](https://leetcode.com/problems/continuous-subarray-sum/) (Medium)
- [974. Subarray Sums Divisible by K](https://leetcode.com/problems/subarray-sums-divisible-by-k/) (Medium)

### What a Google Interviewer Would Ask Next
```
1. "What if we also need to return the actual subarrays, not just the count?"
   → Store (start, end) indices alongside prefix sums in the map
   → Trade-off: O(n) extra space for indices

2. "What if the input is a stream of numbers instead of a fixed array?"
   → Maintain a running prefix sum and hash map
   → Works the same way, just process one element at a time

3. "Can you solve this for a 2D matrix (submatrix sum equals k)?"
   → Fix two rows, compute column prefix sums to reduce to 1D problem
   → O(n^2 * m) time using this technique per row pair

4. "What if we want the longest subarray with sum k instead of the count?"
   → Store only the first occurrence index of each prefix sum
   → Track max(j - i) where prefix[j] - prefix[i] == k

5. "How would you handle this with very large distributed data?"
   → Partition the array, compute local prefix sums per partition
   → Merge boundary prefix sums across partitions
```
