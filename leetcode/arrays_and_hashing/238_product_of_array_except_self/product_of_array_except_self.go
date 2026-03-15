package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	leftProduct := 1
	for i, num := range nums {
		answer[i] = leftProduct
		leftProduct *= num
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}
