package LinkedList

func SortList(head *ListNode) *ListNode {
	return sortList(head)
}

func sortList(head *ListNode) *ListNode {
	// Merge Sort
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	a := head
	b := head.Next

	for b != nil {

		b = b.Next

		if b != nil {
			b = b.Next
			a = a.Next
		}

		if b == nil {
			temp := a.Next
			a.Next = nil
			a = temp
		}
	}

	head = sortList(head)
	a = sortList(a)

	return mergeList(head, a)
}

func mergeList(a, b *ListNode) *ListNode {
	pre, c := &ListNode{}, &ListNode{}
	pre.Next = c
	for a != nil && b != nil {
		if a.Val > b.Val {
			c.Next = b
			b = b.Next
		} else {
			c.Next = a
			a = a.Next
		}

		c = c.Next
	}

	if a != nil {
		c.Next = a
	} else {
		c.Next = b
	}
	return pre.Next.Next
}
