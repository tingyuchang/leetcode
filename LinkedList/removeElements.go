package LinkedList

func RemoveElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	node := dummyNode
	for node != nil {
		if node.Next != nil && node.Next.Val == val {
			if node.Next.Next == nil {
				node.Next = nil
				break
			} else {
				node.Next = node.Next.Next
				continue
			}
		}
		node = node.Next
	}

	return dummyNode.Next
}
