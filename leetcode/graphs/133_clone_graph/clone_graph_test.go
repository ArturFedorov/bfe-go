package clone_graph

import "testing"

func TestCloneGraph(t *testing.T) {
	t.Run("clone 4-node graph", func(t *testing.T) {
		node1 := &Node{Val: 1}
		node2 := &Node{Val: 2}
		node3 := &Node{Val: 3}
		node4 := &Node{Val: 4}
		node1.Neighbors = []*Node{node2, node4}
		node2.Neighbors = []*Node{node1, node3}
		node3.Neighbors = []*Node{node2, node4}
		node4.Neighbors = []*Node{node1, node3}

		clone := cloneGraph(node1)
		if clone == nil {
			t.Fatal("cloneGraph() returned nil, want non-nil")
		}
		if clone == node1 {
			t.Error("cloneGraph() returned same reference, want deep copy")
		}
		if clone.Val != 1 {
			t.Errorf("cloneGraph().Val = %v, want 1", clone.Val)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		got := cloneGraph(nil)
		if got != nil {
			t.Errorf("cloneGraph(nil) = %v, want nil", got)
		}
	})

	t.Run("single node no neighbors", func(t *testing.T) {
		node := &Node{Val: 1, Neighbors: []*Node{}}
		clone := cloneGraph(node)
		if clone == nil {
			t.Fatal("cloneGraph() returned nil, want non-nil")
		}
		if clone == node {
			t.Error("cloneGraph() returned same reference, want deep copy")
		}
		if clone.Val != 1 {
			t.Errorf("clone.Val = %v, want 1", clone.Val)
		}
		if len(clone.Neighbors) != 0 {
			t.Errorf("len(clone.Neighbors) = %v, want 0", len(clone.Neighbors))
		}
	})
}
