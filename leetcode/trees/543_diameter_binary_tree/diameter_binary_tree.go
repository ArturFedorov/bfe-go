package diameter_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		if leftHeight+rightHeight > diameter {
			diameter = leftHeight + rightHeight
		}

		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)

	return diameter
}
