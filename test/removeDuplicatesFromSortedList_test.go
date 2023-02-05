package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/removeDuplicatesfromSortedList"
	"testing"
)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	var testData = []struct {
		input    *removeDuplicatesfromSortedList.ListNode
		expected string
	}{
		{generateListNode([]int{1, 1, 2}), "12"},
		{generateListNode([]int{1, 1, 2, 3, 3}), "123"},
		{generateListNode([]int{0, 0, 0, 0, 0}), "0"},
	}

	for _, tt := range testData {
		result := removeDuplicatesfromSortedList.DeleteDuplicates(tt.input)
		assert.Equal(t, result.Show(), tt.expected)
	}
}

func generateListNode(nums []int) *removeDuplicatesfromSortedList.ListNode {
	head, node := &removeDuplicatesfromSortedList.ListNode{nums[0], nil}, &removeDuplicatesfromSortedList.ListNode{nums[0], nil}
	node = head
	for i := 1; i < len(nums); i++ {
		next := &removeDuplicatesfromSortedList.ListNode{nums[i], nil}
		node.Next = next
		node = next
	}
	return head
}
