package two_sum

func twoSum(nums []int, target int) []int {
	lookUp := make(map[int]int, len(nums))

	for i, num := range nums {
		opposite, ok := lookUp[target-num]
		if ok {
			return []int{opposite, i}
		}
		lookUp[num] = i
	}

	return []int{}
}
