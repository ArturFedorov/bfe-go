package lca_binary_tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2, Left: node7, Right: node4}
	node6 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5, Left: node6, Right: node2}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node1 := &TreeNode{Val: 1, Left: node0, Right: node8}
	root := &TreeNode{Val: 3, Left: node5, Right: node1}

	tests := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want int
	}{
		{
			name: "p=5, q=1 -> LCA is 3",
			root: root,
			p:    node5,
			q:    node1,
			want: 3,
		},
		{
			name: "p=5, q=4 -> LCA is 5",
			root: root,
			p:    node5,
			q:    node4,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if got == nil || got.Val != tt.want {
				gotVal := -1
				if got != nil {
					gotVal = got.Val
				}
				t.Errorf("lowestCommonAncestor() = %v, want %v", gotVal, tt.want)
			}
		})
	}
}
