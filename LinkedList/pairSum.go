package LinkedList

func pairSum(head *ListNode) int {
	slow, fast, next := head, head, head.Next
	fast = fast.Next
	for fast.Next != nil {
		fast = fast.Next.Next
		temp := next.Next
		next.Next = slow
		slow = next
		next = temp

	}

	head.Next = nil
	ans := 0
	for next != nil {
		sum := next.Val + slow.Val

		if sum > ans {
			ans = sum
		}

		next = next.Next
		slow = slow.Next
	}

	return ans
}

/*

approach1: using array to store value in linked list

and iterating array to find maximum


approach2: using fast, slow to get middle og linked list

and reverse 1st half of linked list

so now fast is the end of list

slow is middle of list

*/
