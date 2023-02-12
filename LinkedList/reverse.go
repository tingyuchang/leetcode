package LinkedList

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	next := node.Next
	for next != nil {
		temp := next.Next
		next.Next = node
		node = next
		next = temp
	}

	head.Next = nil
	return node
}
