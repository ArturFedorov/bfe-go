package subarray_sums_divisible_by_k

func subarraysDivByK(nums []int, k int) int {
	count := 0
	prefixMod := 0
	freq := make(map[int]int)
	freq[0] = 1

	for _, n := range nums {
		prefixMod = ((prefixMod+n)%k + k) % k
		count += freq[prefixMod]
		freq[prefixMod]++
	}

	return count
}
