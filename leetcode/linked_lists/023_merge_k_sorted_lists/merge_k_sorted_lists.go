package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil || list2 != nil {
		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]

		merged := mergeTwoLists(l1, l2)
		lists = append(lists, merged)
	}

	return lists[0]
}
