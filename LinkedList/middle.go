package LinkedList

func MiddleNode(head *ListNode) *ListNode {
	return middleNode(head)
}

func middleNode(head *ListNode) *ListNode {
	node := head
	slow := &ListNode{}
	slow.Next = node
	count := 0

	for node != nil {
		slow = slow.Next

		node = node.Next
		if node != nil {
			node = node.Next
			count++
		}

		count++
	}
	if count%2 == 0 {
		return slow.Next
	}
	return slow
}
