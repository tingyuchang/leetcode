package LinkedList

func OddEvenList(head *ListNode) *ListNode {
	return oddEvenList(head)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	first := head
	second := head.Next
	preSecond := &ListNode{}
	preSecond.Next = second

	for first != nil && second != nil {
		first.Next = second.Next
		// check first is nil
		if first.Next == nil {
			second.Next = nil
			break
		}
		first = first.Next
		second.Next = first.Next
		second = second.Next
	}

	// first & second both nil

	first.Next = preSecond.Next

	return head
}
