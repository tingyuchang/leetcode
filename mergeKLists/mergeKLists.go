package mergeKLists

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)

	var tempNode *ListNode
	if k >=1 {
		tempNode = lists[0]
	}
	for k >= 2 {
		tempNode = mergeTwoLists(tempNode, lists[k-1])
		k--
	}

	return tempNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	tempNode := dummyNode
	for l1 != nil || l2 != nil {
		if l1 == nil  {
			tempNode.Next = l2
			break
		}

		if l2 == nil {
			tempNode.Next = l1
			break
		}

		if l1.Val < l2.Val {
			tempNode.Next = l1
			l1 = l1.Next
		} else {
			tempNode.Next = l2
			l2 = l2.Next
		}
		tempNode = tempNode.Next
	}

	return dummyNode.Next
}