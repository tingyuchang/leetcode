package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/mergeTwoList"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	testData := []struct{
		input *mergeTwoList.ListNode
		input2 *mergeTwoList.ListNode
		expected *mergeTwoList.ListNode
	}{
		{	mergeNodeGenerator([]int{1,2,4}),
			mergeNodeGenerator([]int{1,3,4}),
			mergeNodeGenerator([]int{1,1,2,3,4,4})},
	}

	for _,td := range testData {
		result := mergeTwoList.MergeTwoLists(td.input, td.input2)

		for td.expected != nil {
			assert.Equal(t, result.Val, td.expected.Val)
			result = result.Next
			td.expected = td.expected.Next
		}
	}
}

func mergeNodeGenerator(si []int) *mergeTwoList.ListNode {
	firstNode := mergeTwoList.ListNode{}
	tempNode := &firstNode
	for i:= 0; i < len(si); i++ {
		tempNode.Next = &mergeTwoList.ListNode{Val: si[i]}
		tempNode = tempNode.Next
	}

	return firstNode.Next
}