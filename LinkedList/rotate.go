package LinkedList

func RotateRight(head *ListNode, k int) *ListNode {
	return rotateRightII(head, k)
}

func rotateRightII(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	cache := make(map[int]*ListNode)

	node := head
	n := 0
	for node != nil {
		cache[n] = node
		node = node.Next
		n++
	}

	k = k % len(cache)

	if k == 0 {
		return head
	}

	cache[len(cache)-k-1].Next = nil

	cache[len(cache)-1].Next = head

	return cache[len(cache)-k]
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}
	slow := head
	fast := head

	for i := 0; i < k; i++ {
		if fast.Next == nil {
			fast = head
			continue
		}
		fast = fast.Next
	}

	if fast == head {
		return head
	}

	for fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	temp := slow.Next
	fast.Next = head
	slow.Next = nil

	return temp
}
