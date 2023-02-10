package LinkedList

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Show() string {
	dummyNode := &ListNode{}
	var s string
	dummyNode.Next = n
	for dummyNode.Next != nil {
		s = s + strconv.Itoa(dummyNode.Next.Val)
		dummyNode = dummyNode.Next
	}
	return s
}

func GenerateNodeFromArray(nums []int) *ListNode {
	head := &ListNode{}
	node := head

	for i := 0; i < len(nums); i++ {
		node.Val = nums[i]
		if i == len(nums)-1 {
			break
		}
		next := &ListNode{}
		node.Next = next
		node = node.Next
	}

	return head
}
