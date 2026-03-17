package longest_palindromic_substring

func longestPalindrome(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		even := expand(s, i, i)
		odd := expand(s, i, i+1)

		if len(result) < len(even) {
			result = even
		}

		if len(result) < len(odd) {
			result = odd
		}
	}

	return result
}

func expand(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}

	return str[left+1 : right]
}
