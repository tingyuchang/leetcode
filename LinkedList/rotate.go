package LinkedList

func RotateRight(head *ListNode, k int) *ListNode {
	return rotateRight2(head, k)
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	first := head
	count := 0

	for first != nil {
		first = first.Next
		count += 1
	}

	k = k % count

	if k == 0 {
		return head
	}

	first = head
	move := k
	for move > 0 {
		first = first.Next
		move -= 1
	}

	preEnd := head
	for first.Next != nil {
		first = first.Next
		preEnd = preEnd.Next
	}

	end := preEnd.Next
	preEnd.Next = nil
	first.Next = head

	return end
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
