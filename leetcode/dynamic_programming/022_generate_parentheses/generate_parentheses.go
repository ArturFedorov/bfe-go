package generate_parentheses

func generateParenthesis(n int) []string {
	combinations := make([]string, 0)

	if n == 0 {
		return combinations
	}

	number := n

	backtrack("", 0, 0, &combinations, number)

	return combinations

}

func backtrack(current string, open, close int, combinations *[]string, number int) {
	if len(current) == number*2 {
		*combinations = append(*combinations, current)
		return
	}

	if open < number {
		backtrack(current+"(", open+1, close, combinations, number)
	}

	if close < open {
		backtrack(current+")", open, close+1, combinations, number)
	}
}
