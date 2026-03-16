package linked_list_cycle

import "testing"

func buildCycleList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		pos      int
		expected bool
	}{
		{"Example1_CycleAtPos1", []int{3, 2, 0, -4}, 1, true},
		{"Example2_CycleAtPos0", []int{1, 2}, 0, true},
		{"Example3_NoCycle", []int{1}, -1, false},
		{"EmptyList", nil, -1, false},
		{"TwoNodes_NoCycle", []int{1, 2}, -1, false},
		{"SingleNode_SelfCycle", []int{1}, 0, true},
		{"LongList_NoCycle", []int{1, 2, 3, 4, 5}, -1, false},
		{"LongList_CycleAtTail", []int{1, 2, 3, 4, 5}, 4, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildCycleList(tt.vals, tt.pos)
			result := hasCycle(head)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
