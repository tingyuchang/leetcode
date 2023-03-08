package LinkedList

/*
0 -> n -> 1 -> n-1 -> 2 -> n-2 -> 3 -> n-3 ...

*/
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	q := make([]*ListNode, 0)
	temp := head
	count := 0
	for temp != nil {
		count++
		q = append(q, temp)
		temp = temp.Next
	}

	node := head

	for i := 1; i <= count/2; i++ {
		temp = node.Next
		next := q[len(q)-1]
		q = q[:len(q)-1]
		node.Next = next
		next.Next = temp
		node = temp
	}
	node.Next = nil

}
