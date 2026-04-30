package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)
	for _, x := range nums {
		set[x] = true
	}

	longest := 1

	for x := range set {
		if !set[x-1] {
			curr := x
			length := 1

			for set[curr+1] {
				curr++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
