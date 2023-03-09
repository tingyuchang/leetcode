package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next

	for slow != nil && fast != nil {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next

		if fast != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}

	if slow == nil || fast == nil {
		return nil
	}

	temp := make(map[*ListNode]int)
	temp[slow] = 0
	count := 1
	slow = slow.Next

	for slow != fast {
		temp[slow] = count
		count++
		slow = slow.Next
	}

	slow = head
	for {
		if _, ok := temp[slow]; ok {
			return slow
		}
		slow = slow.Next
	}
}

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
