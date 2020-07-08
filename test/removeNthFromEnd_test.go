package test

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"leetcode/removeNthFromEnd"
	)


func TestRemoveNthFromEnd(t *testing.T) {
	testData := []struct{
		input *removeNthFromEnd.ListNode
		target int
		expected *removeNthFromEnd.ListNode
	}{
		{nodeGenerator([]int{1,2,3,4,5}), 2, nodeGenerator([]int{1,2,3,5})},
		{nodeGenerator([]int{1,2}), 1, nodeGenerator([]int{1})},
		{nodeGenerator([]int{1,2}), 2, nodeGenerator([]int{2})},

	}

	for _,td := range testData {
		result := removeNthFromEnd.RemoveNthFromEnd(td.input, td.target)
		for result.Next != nil {
			assert.Equal(t, result.Val, td.expected.Val)
			result = result.Next
			td.expected = td.expected.Next
		}
	}

}


func nodeGenerator(si []int) *removeNthFromEnd.ListNode {
	firstNode := removeNthFromEnd.ListNode{0,nil}
	tempNode := &firstNode
	for i:= 0; i < len(si); i++ {
		tempNode.Next = &removeNthFromEnd.ListNode{si[i], nil}
		tempNode = tempNode.Next
	}

	return firstNode.Next
}