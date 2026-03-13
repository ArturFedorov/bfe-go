package diameter_binary_tree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "[1,2,3,4,5] -> 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
			want: 3,
		},
		{
			name: "[1,2] -> 1",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			want: 1,
		},
		{
			name: "nil -> 0",
			root: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diameterOfBinaryTree(tt.root)
			if got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
