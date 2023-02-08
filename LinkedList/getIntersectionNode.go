package LinkedList

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}

func GetIntersectionNodeError(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}

	nodeA, nodeB := headA, headB

	for nodeA.Next != nil {
		nodeA = nodeA.Next
	}
	for nodeB.Next != nil {
		nodeB = nodeB.Next
	}

	if nodeA != nodeB {
		return nil
	}

	// there exist intersection

	nodeA, nodeB = headA, headB.Next

	if nodeA == nil {
		nodeA = headA
	}

	if nodeB == nil {
		nodeB = headB
	}

	// prepare previous node to help
	var preA, preB *ListNode
	preA = nil
	preB = headB
	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB && preA != preB {
			return nodeA
		}

		preA = nodeA
		nodeA = nodeA.Next
		nodeB = nodeB.Next

		if nodeA == nil {
			preA = nil
			nodeA = headA
		}

		if nodeB == nil {
			preB = nil
			nodeB = headB
		}

		if nodeB.Next != nil {
			preB = nodeB
			nodeB = nodeB.Next
		} else {
			preB = nil
			nodeB = headB
		}

		if nodeA == headA && nodeB == headB.Next {
			preB = headB
			nodeB = nodeB.Next.Next
		}
	}

	return nil
}
