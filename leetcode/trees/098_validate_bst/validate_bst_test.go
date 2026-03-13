package validate_bst

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "valid BST [2,1,3]",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "invalid BST [5,1,4,nil,nil,3,6]",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "nil tree",
			root: nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.root)
			if got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
