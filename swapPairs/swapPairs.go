package swapPairs

import (
	"leetcode/models"
)

func SwapPairs(head *models.ListNode) *models.ListNode {
	firstNode := &models.ListNode{}
	firstNode.Next = head
	preNode := firstNode

	for head != nil && head.Next != nil {
		secondNode := head.Next
		head.Next = secondNode.Next
		secondNode.Next = head
		preNode.Next = secondNode
		preNode = head
		head = head.Next
	}
	return firstNode.Next
}












/*

firstNode := &models.ListNode{}
	firstNode.Next = head
	preNode := firstNode
	fmt.Printf("00. head: %v|%p\tpreNode: %v|%p\tfirst: %v|%p\n", head, head, preNode, preNode, firstNode, firstNode)

	for (head != nil) && head.Next != nil {
		secNode := head.Next
		fmt.Printf("1. secNode: %v\n", secNode)
		head.Next = secNode.Next
		fmt.Printf("2. head: %v\n", head)
		secNode.Next = head
		fmt.Printf("3. secNode: %v\n", secNode)
		preNode.Next = secNode
		fmt.Printf("4. perNode: %v\n", preNode)
		preNode = head
		fmt.Printf("5. perNode: %v\n", preNode)
		head = head.Next
		fmt.Printf("6. head: %v\n", head)
	}
	fmt.Printf("00. head: %v|%p\tpreNode: %v|%p\tfirst: %v|%p\n", head, head, preNode, preNode, firstNode, firstNode)
	return firstNode.Next




 */