package invert_binary_tree

import "testing"

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "[4,2,7,1,3,6,9] -> [4,7,2,9,6,3,1]",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "nil -> nil",
			root: nil,
			want: nil,
		},
		{
			name: "[2,1,3] -> [2,3,1]",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.root)
			gotSlice := treeToSlice(got)
			if !sliceEqual(gotSlice, tt.want) {
				t.Errorf("invertTree() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
