package removeDuplicatesfromSortedList

import (
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Show() string {
	var s string
	temp := n
	for temp.Next != nil {
		s = s + strconv.Itoa(temp.Val)
		temp = temp.Next
	}
	s = s + strconv.Itoa(temp.Val)
	return s
}

func DeleteDuplicates(head *ListNode) *ListNode {
	dummyNode, node := &ListNode{-101, nil}, &ListNode{}

	dummyNode.Next = head

	for dummyNode.Next != nil {
		if dummyNode.Val != dummyNode.Next.Val {
			node.Next = dummyNode.Next
			node = node.Next
		}

		dummyNode = dummyNode.Next
	}

	if dummyNode.Val == node.Val {
		node.Next = nil
	}

	return head
}

func DeleteDuplicatesV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			if node.Next.Next == nil {
				node.Next = nil
				return head
			} else {
				node.Next = node.Next.Next
				continue
			}
		}
		node = node.Next
	}

	return head
}
