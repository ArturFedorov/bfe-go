# 974. Subarray Sums Divisible by K

**Difficulty:** Medium

**Topics:** Array, Hash Table, Prefix Sum

---

## Description

Given an integer array `nums` and an integer `k`, return _the number of non-empty subarrays that have a sum divisible by_ `k`.

A subarray is a contiguous part of an array.

---

## Examples

### Example 1

```
Input: nums = [4,5,0,-2,-3,1], k = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by k = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
```

### Example 2

```
Input: nums = [5], k = 9
Output: 0
```

---

## Constraints

- `1 <= nums.length <= 3 * 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `2 <= k <= 10^4`

---

## Approach Hints

<details>
<summary>Hint 1</summary>
This is very similar to "Subarray Sum Equals K" (LC 560), but instead of exact sum you care about divisibility.
</details>

<details>
<summary>Hint 2</summary>
If prefix[j] % k == prefix[i] % k, then the subarray from i+1 to j has a sum divisible by k.
</details>

<details>
<summary>Hint 3</summary>
Use a hash map to count prefix sum remainders. Be careful with negative remainders — in Go, use ((sum % k) + k) % k to normalize.
</details>

---

## Related Problems

- [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/) (Medium)
- [523. Continuous Subarray Sum](https://leetcode.com/problems/continuous-subarray-sum/) (Medium)
- [1590. Make Sum Divisible by P](https://leetcode.com/problems/make-sum-divisible-by-p/) (Medium)

### What a Google Interviewer Would Ask Next
```
1. "What if k can be negative?"
   → Take abs(k) since divisibility is the same for k and -k

2. "How would you return the actual subarrays, not just the count?"
   → Store indices for each remainder in the map
   → For each remainder match, enumerate all (i, j) pairs

3. "What if the array is a stream and you need a running count?"
   → Maintain running prefix sum mod k and the remainder frequency map
   → Works identically in streaming fashion

4. "How does this relate to LC 560 (Subarray Sum Equals K)?"
   → LC 560 checks prefix[j] - prefix[i] == k (exact match)
   → This checks prefix[j] % k == prefix[i] % k (same remainder)
   → Both use prefix sum + hash map, O(n) time O(n) space

5. "Can you solve it in O(1) space?"
   → Only if k is small: use an array of size k instead of a hash map
   → Still O(k) space technically, but avoids hash overhead
```
