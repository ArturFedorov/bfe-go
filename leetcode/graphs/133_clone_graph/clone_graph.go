package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	copies := make([]*Node, 101)
	dfs(node, copies)

	return copies[node.Val]
}

func dfs(node *Node, copies []*Node) {
	newNode := new(Node)
	newNode.Val = node.Val

	copies[node.Val] = newNode

	for _, neighbour := range node.Neighbors {
		if copies[neighbour.Val] == nil {
			dfs(neighbour, copies)
		}

		newNode.Neighbors = append(newNode.Neighbors, copies[neighbour.Val])
	}
}
