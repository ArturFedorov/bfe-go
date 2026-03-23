package lca_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var foundP, foundQ bool
	result := dfs(root, p, q, &foundP, &foundQ)
	if foundP && foundQ {
		return result
	}
	return nil // one or both don't exist
}

func dfs(root, p, q *TreeNode, foundP, foundQ *bool) *TreeNode {
	if root == nil {
		return nil
	}

	left := dfs(root.Left, p, q, foundP, foundQ)
	right := dfs(root.Right, p, q, foundP, foundQ)

	// check AFTER recursing so we don't short-circuit
	if root == p {
		*foundP = true
		return root
	}
	if root == q {
		*foundQ = true
		return root
	}

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
