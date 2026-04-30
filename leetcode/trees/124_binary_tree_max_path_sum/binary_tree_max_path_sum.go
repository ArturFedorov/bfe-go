package binary_tree_max_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMaxPath := max(helper(node.Left), 0)
		rightMaxPath := max(helper(node.Right), 0)

		maxIfNodeIsRoot := node.Val + leftMaxPath + rightMaxPath
		maxSum = max(maxSum, maxIfNodeIsRoot)

		return node.Val + max(leftMaxPath, rightMaxPath)
	}

	helper(root)

	return maxSum
}
