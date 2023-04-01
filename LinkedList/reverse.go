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
