package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	prefixSumFrequencies := make(map[int]int)
	prefixSumFrequencies[0] = 1

	result := 0
	prefixSum := 0

	for _, num := range nums {
		prefixSum += num
		diff := prefixSum - k
		result += prefixSumFrequencies[diff]
		prefixSumFrequencies[prefixSum]++
	}

	return result
}
