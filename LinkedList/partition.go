package LinkedList

func Partition(head *ListNode, x int) *ListNode {
	return partition(head, x)
}

func partition(head *ListNode, x int) *ListNode {
	dummyNode, dummyNode2, greaterNode, lessNode := &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}
	node := head

	dummyNode.Next = lessNode
	dummyNode2.Next = greaterNode

	for node != nil {
		if node.Val < x {
			// add to less node
			lessNode.Next = node
			lessNode = lessNode.Next
		} else {
			// add to greater node
			greaterNode.Next = node
			greaterNode = greaterNode.Next
		}

		node = node.Next
		lessNode.Next = nil
		greaterNode.Next = nil
	}
	lessNode.Next = dummyNode2.Next.Next

	return dummyNode.Next.Next
}
