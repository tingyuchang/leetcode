package LinkedList

func SwapPairs(head *ListNode) *ListNode {
	firstNode := &ListNode{}
	firstNode.Next = head
	preNode := firstNode

	for head != nil && head.Next != nil {
		secondNode := head.Next
		head.Next = secondNode.Next
		secondNode.Next = head
		preNode.Next = secondNode
		preNode = head
		head = head.Next
	}
	return firstNode.Next
}
