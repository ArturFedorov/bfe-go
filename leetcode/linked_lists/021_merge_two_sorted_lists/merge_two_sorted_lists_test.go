package merge_two_sorted_lists

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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{"Example1", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"Example2", nil, nil, nil},
		{"Example3", nil, []int{0}, []int{0}},
		{"OneEmpty", []int{1, 2, 3}, nil, []int{1, 2, 3}},
		{"AllSame", []int{1, 1, 1}, []int{1, 1}, []int{1, 1, 1, 1, 1}},
		{"NoOverlap", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"SingleElements", []int{2}, []int{1}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToList(tt.list1)
			l2 := sliceToList(tt.list2)
			result := listToSlice(mergeTwoLists(l1, l2))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
