package reverse_nodes_in_k_group

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{"Example1", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"Example2", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{"KEqualsOne", []int{1, 2, 3}, 1, []int{1, 2, 3}},
		{"KEqualsLength", []int{1, 2, 3}, 3, []int{3, 2, 1}},
		{"KGreaterThanLength", []int{1, 2}, 3, []int{1, 2}},
		{"SingleNode", []int{1}, 1, []int{1}},
		{"TwoNodes_K2", []int{1, 2}, 2, []int{2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := listToSlice(reverseKGroup(head, tt.k))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
