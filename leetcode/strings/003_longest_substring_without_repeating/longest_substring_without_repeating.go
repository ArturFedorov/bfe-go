package longest_substring_without_repeating

func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int)
	left := 0
	result := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq[s[right]] > 1 {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}

	return result
}
