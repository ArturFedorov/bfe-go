package validate_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	minValue := math.MinInt64
	isValid := true

	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}

		visit(node.Left)

		if node.Val <= minValue {
			isValid = false
		}
		minValue = node.Val

		visit(node.Right)
	}

	visit(root)

	return isValid
}
