package serialize_deserialize

import "testing"

func treesEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treesEqual(a.Left, b.Left) && treesEqual(a.Right, b.Right)
}

func TestCodec(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "[1,2,3,nil,nil,4,5] roundtrip",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
		{
			name: "nil roundtrip",
			root: nil,
		},
		{
			name: "single node roundtrip",
			root: &TreeNode{Val: 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Codec{}
			data := c.serialize(tt.root)
			got := c.deserialize(data)
			if !treesEqual(got, tt.root) {
				t.Errorf("roundtrip failed: serialize then deserialize did not reproduce original tree")
			}
		})
	}
}
