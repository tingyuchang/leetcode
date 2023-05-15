package LinkedList

func SwapNodes(head *ListNode, k int) *ListNode {
	return swapNodes(head, k)
}

func swapNodes(head *ListNode, k int) *ListNode {
	current, first, second := &ListNode{Next: head}, &ListNode{Next: head}, &ListNode{Next: head}

	for i := 0; i < k; i++ {
		first = first.Next
		current = current.Next
	}

	for current != nil {
		current = current.Next
		second = second.Next
	}

	first.Val, second.Val = second.Val, first.Val

	return head
}
