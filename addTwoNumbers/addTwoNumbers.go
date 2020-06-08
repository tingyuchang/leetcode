package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstNode *ListNode
	loopNode1 := &ListNode{}
	loopNode2 := &ListNode{}
	isCarry := false

	for {
		val := l1.Val + l2.Val
		if isCarry {
			val++
		}
		if val >= 10 {
			isCarry = true
			val -= 10
		} else {
			isCarry = false
		}

		if firstNode == nil {
			firstNode = &ListNode{
				Val:  val,
				Next: nil,
			}

			if l1.Next != nil || l2.Next != nil {
				firstNode.Next = loopNode1
				loopNode1.Next = loopNode2
			}
		} else {
			loopNode1.Val = val

			if l1.Next != nil || l2.Next != nil {
				loopNode1 = loopNode2
				loopNode2 = &ListNode{}
				loopNode1.Next = loopNode2
			}

		}

		if l1.Next == nil && l2.Next == nil {

			if isCarry {
				if firstNode.Next == nil {
					firstNode.Next = &ListNode{1, nil}
				} else {
					loopNode1.Next = &ListNode{1, nil}
				}
			} else {
				loopNode1.Next = nil
			}
			break
		} else {

			if l1.Next == nil {
				l1 = &ListNode{0, nil}
			} else {
				l1 = l1.Next
			}

			if l2.Next == nil {
				l2 = &ListNode{0, nil}
			} else {
				l2 = l2.Next
			}

		}
	}

	return firstNode
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	var carry int
	firstNode := ListNode{0,nil}
	tempNode := &firstNode
	for l1 != nil || l2 != nil || carry > 0{
		var sum, v1, v2 int

		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}

		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		sum = v1 + v2 + carry
		carry = sum/10
		sum = sum % 10

		tempNode.Next = &ListNode{sum, nil}
		tempNode = tempNode.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		tempNode.Next = &ListNode{carry, nil}
		carry = 0
	}

	return firstNode.Next
}