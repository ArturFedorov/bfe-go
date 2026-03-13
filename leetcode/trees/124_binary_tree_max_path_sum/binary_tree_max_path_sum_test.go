package binary_tree_max_path_sum

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "[1,2,3] -> 6",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: 6,
		},
		{
			name: "[-10,9,20,nil,nil,15,7] -> 42",
			root: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 42,
		},
		{
			name: "single negative node [-3]",
			root: &TreeNode{Val: -3},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxPathSum(tt.root)
			if got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
