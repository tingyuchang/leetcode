package LinkedList

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)

	var tmpNode *ListNode
	if k >= 1 {
		tmpNode = lists[0]
	}
	for k >= 2 {
		tmpNode = mergeTwoLists(tmpNode, lists[k-1])
		k--
	}

	return tmpNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode
	node := &ListNode{}

	if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		head = list1
		list1 = list1.Next
	}
	node.Next = head

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			head.Next = list2
			list2 = list2.Next
		} else {
			head.Next = list1
			list1 = list1.Next
		}
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	}

	if list2 != nil {
		head.Next = list2
	}

	return node.Next
}
