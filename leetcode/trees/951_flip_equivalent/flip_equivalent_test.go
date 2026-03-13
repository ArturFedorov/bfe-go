package flip_equivalent

import "testing"

func TestFlipEquiv(t *testing.T) {
	tests := []struct {
		name  string
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		{
			name: "flip equivalent trees",
			root1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 8}}},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			root2: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: true,
		},
		{
			name:  "both nil",
			root1: nil,
			root2: nil,
			want:  true,
		},
		{
			name:  "[0,nil,1] vs [0,1]",
			root1: &TreeNode{Val: 0, Right: &TreeNode{Val: 1}},
			root2: &TreeNode{Val: 0, Left: &TreeNode{Val: 1}},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flipEquiv(tt.root1, tt.root2)
			if got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
