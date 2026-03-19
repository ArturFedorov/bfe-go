package valid_parentheses

func isValid(s string) bool {
	stack := make([]rune, 0)
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if match, found := brackets[char]; found {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
