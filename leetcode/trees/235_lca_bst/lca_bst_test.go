package lca_bst

import "testing"

func TestLowestCommonAncestorBST(t *testing.T) {
	node0 := &TreeNode{Val: 0}
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4, Left: node3, Right: node5}
	node2 := &TreeNode{Val: 2, Left: node0, Right: node4}
	node7 := &TreeNode{Val: 7}
	node9 := &TreeNode{Val: 9}
	node8 := &TreeNode{Val: 8, Left: node7, Right: node9}
	root := &TreeNode{Val: 6, Left: node2, Right: node8}

	tests := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want int
	}{
		{
			name: "p=2, q=8 -> LCA is 6",
			root: root,
			p:    node2,
			q:    node8,
			want: 6,
		},
		{
			name: "p=2, q=4 -> LCA is 2",
			root: root,
			p:    node2,
			q:    node4,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestorBST(tt.root, tt.p, tt.q)
			if got == nil || got.Val != tt.want {
				gotVal := -1
				if got != nil {
					gotVal = got.Val
				}
				t.Errorf("lowestCommonAncestorBST() = %v, want %v", gotVal, tt.want)
			}
		})
	}
}
