package serialize_deserialize

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	var res []string

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			res = append(res, "N")
			return
		}

		res = append(res, strconv.Itoa(node.Val))
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)

	return strings.Join(res, ",")
}

func (c *Codec) deserialize(data string) *TreeNode {
	nodeValues := strings.Split(data, ",")
	index := 0

	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		if nodeValues[index] == "N" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(nodeValues[index])
		index++

		node := &TreeNode{Val: val}

		node.Left = buildTree()
		node.Right = buildTree()

		return node
	}

	return buildTree()
}
