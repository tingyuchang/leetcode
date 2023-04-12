package __InterviewBit

func swapPairs(A *listNode) *listNode {
	if A == nil {
		return nil
	}

	if A.next == nil {
		return A
	}
	temp := &listNode{}
	node := A
	next := A.next

	temp.next = next

	for node != nil && next != nil {
		node.next = next.next
		next.next = node
		temp := node
		node = node.next
		if node != nil {
			next = node.next
			temp.next = next
			if next == nil {
				temp.next = node
				break
			}
		} else {
			break
		}
	}

	return temp.next

}
