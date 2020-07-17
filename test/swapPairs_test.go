package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/models"
	"leetcode/swapPairs"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	testData := []struct{
		input *models.ListNode
		expected *models.ListNode
	}{
		{GenerateListNode([]int{1,2,3,4}), GenerateListNode([]int{2,1,4,3,4})},
	}

	for _,td := range testData {
		result := swapPairs.SwapPairs(td.input)

		for td.expected.Next != nil {
			assert.Equal(t, 2, 4)
			result = result.Next
			td.expected = td.expected.Next
		}
	}
}

func GenerateListNode(si []int) *models.ListNode {
	firstNode := &models.ListNode{}
	tempNode := firstNode
	for _,v := range si {
		tempNode.Next = &models.ListNode{Val: v}
		tempNode = tempNode.Next
	}


	return firstNode.Next
}