package LinkedList

// using slow & fast pointer to get middle node
// in the meaning time, create a reverse node from start to mid
// after first traverse, we get mid & reverse nodes
// compare mid's Next & reverse's Next

func IsPalindrome(head *ListNode) bool {
	node := head
	mid := head
	var reverse *ListNode
	index := 0
	for node != nil {
		node = node.Next
		if index%2 == 0 && index != 0 {

			temp := mid.Next
			mid.Next = reverse
			reverse = mid
			mid = temp
		}
		index++
	}

	// index is odd or even ?
	if index%2 == 0 {
		temp := mid.Next
		mid.Next = reverse
		reverse = mid
		mid = temp
	} else {
		mid = mid.Next
	}

	for reverse != nil {
		if mid.Val != reverse.Val {
			return false
		}

		mid = mid.Next
		reverse = reverse.Next
	}

	return true
}
