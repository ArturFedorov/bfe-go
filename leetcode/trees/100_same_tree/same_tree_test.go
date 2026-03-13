package same_tree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "same trees [1,2,3]",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "different structure [1,2] vs [1,nil,2]",
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "both nil",
			p:    nil,
			q:    nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
