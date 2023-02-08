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
