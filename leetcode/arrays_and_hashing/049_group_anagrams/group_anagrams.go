package group_anagrams

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}

		for _, r := range str {
			count[r-'a']++
		}

		groups[count] = append(groups[count], str)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
