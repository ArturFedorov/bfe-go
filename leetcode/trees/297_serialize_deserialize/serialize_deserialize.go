package serialize_deserialize

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func (c *Codec) serialize(root *TreeNode) string {
	return ""
}

func (c *Codec) deserialize(data string) *TreeNode {
	return nil
}
