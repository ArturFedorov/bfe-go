package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	n := len(nums)
	low, high, ans := 0, n-1, -1

	right := nums[n-1]

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] <= right {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return nums[ans]
}
