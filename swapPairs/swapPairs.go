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