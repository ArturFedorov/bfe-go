package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// counting elements
	var count int
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	// reverse
	zeroNode := &ListNode{Next: head}
	prev := zeroNode

	for i := 0; i < count/k; i++ {
		curr = prev.Next
		for j := 1; j < k; j++ {
			nextNode := curr.Next
			curr.Next = nextNode.Next
			nextNode.Next = prev.Next
			prev.Next = nextNode
		}
		prev = curr
	}

	return zeroNode.Next
}
