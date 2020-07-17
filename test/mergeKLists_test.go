package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/mergeKLists"
	"testing"
)

func TestMergeKList(t *testing.T) {
	testData := []struct{
		input []*mergeKLists.ListNode
		expected *mergeKLists.ListNode
	}{
		{[]*mergeKLists.ListNode{
			generateList([]int{1,4,5}),
			generateList([]int{1,3,4}),
			generateList([]int{2,6}),
		}, generateList([]int{1,1,2,3,4,4,5,6})},
		{[]*mergeKLists.ListNode{
			generateList([]int{1}),
		}, generateList([]int{1})},
	}

	for _,td := range testData {
		result := mergeKLists.MergeKLists(td.input)

		for td.expected != nil {
			assert.Equal(t, result.Val, td.expected.Val)
			td.expected = td.expected.Next
			result = result.Next
		}
	}
}

func generateList(firstNodes[]int) (*mergeKLists.ListNode) {
	l1 := mergeKLists.ListNode{}
	tempNode1 := &l1

	for _,v := range firstNodes {
		tempNode1.Next = &mergeKLists.ListNode{Val: v}
		tempNode1 = tempNode1.Next
	}


	return l1.Next

}