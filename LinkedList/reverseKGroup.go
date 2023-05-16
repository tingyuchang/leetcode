package LinkedList

func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	available := k
	node := head
	for node != nil && available > 0 {
		node = node.Next
		available -= 1
	}

	if available == 0 {
		var current, pre, next *ListNode
		current = head
		count := 0

		for count < k && current != nil {
			next = current.Next
			current.Next = pre
			pre = current
			current = next
			count += 1
		}
		if next != nil {
			head.Next = reverseKGroup(next, k)
		}
		return pre
	}

	return head
}
