package binary_tree_right_side_view

import "testing"

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "[1,2,3,nil,5,nil,4] -> [1,3,4]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "[1,nil,3] -> [1,3]",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			want: []int{1, 3},
		},
		{
			name: "nil -> []",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rightSideView(tt.root)
			if got == nil {
				got = []int{}
			}
			if len(got) != len(tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("rightSideView() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
