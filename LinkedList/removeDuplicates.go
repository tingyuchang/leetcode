package LinkedList

func DeleteDuplicatesII(head *ListNode) *ListNode {
	return deleteDuplicatesII(head)
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummyNode, temp := &ListNode{}, &ListNode{}
	temp.Next = dummyNode
	node := head
	isDuplicated := false
	for node != nil {
		if node.Next != nil && node.Next.Val == node.Val {
			isDuplicated = true
			if node.Next.Next == nil {
				node.Next = nil
				dummyNode.Next = nil
				return temp.Next.Next
			} else {
				if node.Next.Next.Val != node.Val {
					isDuplicated = false
				}
				node = node.Next.Next
				continue
			}
		}
		if !isDuplicated {
			dummyNode.Next = node
			dummyNode = dummyNode.Next
		} else {
			if node.Next == nil {
				dummyNode.Next = nil
			}
		}

		node = node.Next
		isDuplicated = false
	}

	return temp.Next.Next
}

func DeleteDuplicates(head *ListNode) *ListNode {
	dummyNode, node := &ListNode{-101, nil}, &ListNode{}

	dummyNode.Next = head

	for dummyNode.Next != nil {
		if dummyNode.Val != dummyNode.Next.Val {
			node.Next = dummyNode.Next
			node = node.Next
		}

		dummyNode = dummyNode.Next
	}

	if dummyNode.Val == node.Val {
		node.Next = nil
	}

	return head
}

func DeleteDuplicatesV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			if node.Next.Next == nil {
				node.Next = nil
				return head
			} else {
				node.Next = node.Next.Next
				continue
			}
		}
		node = node.Next
	}

	return head
}
