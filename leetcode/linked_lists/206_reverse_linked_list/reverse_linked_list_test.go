package reverse_linked_list

import (
	"reflect"
	"testing"
)

func sliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Example1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"Example2", []int{1, 2}, []int{2, 1}},
		{"SingleNode", []int{1}, []int{1}},
		{"EmptyList", nil, nil},
		{"ThreeElements", []int{1, 2, 3}, []int{3, 2, 1}},
	}

	t.Run("Iterative", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				head := sliceToList(tt.input)
				result := listToSlice(reverseListIterative(head))
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			})
		}
	})

	t.Run("Recursive", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				head := sliceToList(tt.input)
				result := listToSlice(reverseListRecursive(head))
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			})
		}
	})
}
