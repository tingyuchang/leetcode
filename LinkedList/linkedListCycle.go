package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next

		if fast != nil {
			fast = fast.Next
		} else {
			return false
		}
	}

	return false
}

func HasCycleSlow(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	m := make(map[*ListNode]int)
	i := 0

	node := head

	for node != nil {

		_, ok := m[node]
		if ok {
			return true
		}
		m[node] = i
		node = node.Next
		i++
	}

	return false
}
