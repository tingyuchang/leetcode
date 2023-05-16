package LinkedList

func ReverseListII(head *ListNode, left, right int) *ListNode {
	return reverseListII(head, left, right)
}

func reverseListII(head *ListNode, left, right int) *ListNode {
	if head.Next == nil {
		return head
	}
	var leftNode *ListNode
	var rightNode *ListNode
	var next *ListNode
	node := head
	position := 1

	if position == left {
		next = node.Next
		node.Next = nil
		rightNode = node
	}

	for node != nil {
		if next == nil {
			if node.Next != nil && position+1 == left {
				leftNode = node
				next = node.Next.Next
				rightNode = node.Next
			}
			node = node.Next
		} else {

			if position == right {
				if rightNode != nil {
					rightNode.Next = next
				}
				if leftNode != nil {
					leftNode.Next = node
				}

				break
			} else {
				temp := next.Next
				next.Next = node
				node = next
				next = temp

				if next == nil {
					if rightNode != nil {
						rightNode.Next = next
					}
					if leftNode != nil {
						leftNode.Next = node
					}
					break
				}
			}
		}
		position++
	}

	if left == 1 {
		return node
	}

	return head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	preFirst := &ListNode{}
	preFirst.Next = head
	first := head
	leftCount := left
	for leftCount > 1 {
		first = first.Next
		preFirst = preFirst.Next
		leftCount -= 1
	}

	var pre, next, tempHead *ListNode
	tempHead = first
	count := right - left + 1
	for count > 0 && first != nil {
		next = first.Next
		first.Next = pre
		pre = first
		first = next
		count -= 1
	}

	tempHead.Next = next
	preFirst.Next = pre

	return head
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	next := node.Next
	for next != nil {
		temp := next.Next
		next.Next = node
		node = next
		next = temp
	}

	head.Next = nil
	return node
}
