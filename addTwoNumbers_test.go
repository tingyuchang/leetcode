package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var loopNode *ListNode
	data := []int{}

	n1 := &ListNode{
		Val: 2,
	}
	n2 := &ListNode{
		Val: 3,
	}
	n3 := &ListNode{
		Val: 4,
	}

	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{
		Val: 5,
	}
	n5 := &ListNode{
		Val: 6,
	}
	n6 := &ListNode{
		Val: 4,
	}
	n4.Next = n5
	n5.Next = n6

	loopNode = addTwoNumbers2(n1, n4)

	for {
		data = append(data, loopNode.Val)
		if loopNode.Next != nil {
			loopNode = loopNode.Next
		} else {
			break
		}
	}

	assert.Equal(t, data, []int{7, 9, 8})

	n1 = &ListNode{
		Val: 9,
	}
	n2 = &ListNode{
		Val: 1,
	}
	n3 = &ListNode{
		Val: 6,
	}
	n1.Next = n2
	n2.Next = n3
	n4 = &ListNode{
		Val: 0,
	}

	data = []int{}
	loopNode = addTwoNumbers2(n1, n4)
	for {
		data = append(data, loopNode.Val)
		if loopNode.Next != nil {
			loopNode = loopNode.Next
		} else {
			break
		}
	}
	assert.Equal(t, data, []int{9, 1, 6})

	n1 = &ListNode{
		Val: 1,
	}
	n2 = &ListNode{
		Val: 8,
	}
	n1.Next = n2
	n4 = &ListNode{
		Val: 0,
	}
	data = []int{}
	loopNode = addTwoNumbers2(n1, n4)
	for {
		data = append(data, loopNode.Val)
		if loopNode.Next != nil {
			loopNode = loopNode.Next
		} else {
			break
		}
	}
	assert.Equal(t, data, []int{1, 8})
}
