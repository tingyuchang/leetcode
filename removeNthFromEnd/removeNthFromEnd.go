package removeNthFromEnd

import "fmt"

// Definition for singly-linked list.
 type ListNode struct {
 	Val int
 	Next *ListNode
 }

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	count := 1
	node := head.Next
	for node != nil {
		count++
		node = node.Next
	}

	if count == n {
		return head.Next
	}

	node = head
	index := 1
	for node != nil {
		if count - n == index  {
			node.Next = node.Next.Next
			fmt.Printf("node: %v\t head: %v\n", node, head)
			break;
		}
		index++
		node = node.Next
	}
	return head
}

