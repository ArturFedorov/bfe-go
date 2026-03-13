package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
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

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "Example1",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "Example2",
			lists: [][]int{},
			want:  nil,
		},
		{
			name:  "Example3",
			lists: [][]int{{}},
			want:  nil,
		},
		{
			name:  "SingleList",
			lists: [][]int{{1, 2, 3}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "TwoLists",
			lists: [][]int{{1, 3}, {2, 4}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "AllEmpty",
			lists: [][]int{{}, {}, {}},
			want:  nil,
		},
		{
			name:  "MixedEmpty",
			lists: [][]int{{}, {1}, {}, {2}},
			want:  []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, s := range tt.lists {
				lists[i] = sliceToList(s)
			}
			got := listToSlice(mergeKLists(lists))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists(%v) = %v, want %v", tt.lists, got, tt.want)
			}
		})
	}
}
