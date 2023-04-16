package LinkedList

import "fmt"

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{}
	node := &ListNode{}

	first.Next = node

	isCarry := false
	for node != nil && l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val

		if isCarry {
			isCarry = false
			sum++
		}

		if sum >= 10 {
			isCarry = true
			sum = sum % 10
		}

		next := &ListNode{
			Val:  sum,
			Next: nil,
		}

		node.Next = next
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	//	1. check l1 / l2 which still available, if no just check carry
	//  2. check carry numbers
	//  3. link node.next to l1/l2
	//

	if l1 != nil {
		node.Next = l1

	} else {
		node.Next = l2
	}

	if node.Next != nil {
		node = node.Next
		for node != nil {
			sum := node.Val

			if isCarry {
				isCarry = false
				sum++
			}

			if sum >= 10 {
				isCarry = true
				sum = sum % 10
			}

			node.Val = sum
			fmt.Println("node:", node.String(), first.Next.Next)
			if node.Next == nil {
				break
			}
			node = node.Next

		}
	}

	if isCarry {
		next := &ListNode{
			Val:  1,
			Next: nil,
		}

		node.Next = next
	}

	return first.Next.Next
}
