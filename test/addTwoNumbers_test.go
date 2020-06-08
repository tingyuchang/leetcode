package test

import (
	"github.com/stretchr/testify/assert"
	"leetcode/addTwoNumbers"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	var testData = []struct {
		input1        	*addTwoNumbers.ListNode
		input2			*addTwoNumbers.ListNode
		expected 		[]int
	}{
		{generateAddTwoNumberstestData([]int{2, 3, 4}), generateAddTwoNumberstestData([]int{5, 6, 4}), []int{7, 9, 8}},
		{generateAddTwoNumberstestData([]int{2, 9, 2}), generateAddTwoNumberstestData([]int{1, 3, 8}), []int{3, 2, 1, 1}},
		{generateAddTwoNumberstestData([]int{9, 1, 6}), generateAddTwoNumberstestData([]int{0}), []int{9, 1, 6}},
		{generateAddTwoNumberstestData([]int{1, 8}), generateAddTwoNumberstestData([]int{0}), []int{1, 8}},
	}

	for _, tt := range testData {
		actual := validateAddTwoNumberstestData(addTwoNumbers.AddTwoNumbers2(tt.input1, tt.input2))
		assert.Equal(t, actual, tt.expected)
	}
}

func generateAddTwoNumberstestData(firstNodes[]int) (*addTwoNumbers.ListNode) {
	l1 := addTwoNumbers.ListNode{}
	tempNode1 := &l1

	for _,v := range firstNodes {
		tempNode1.Next = &addTwoNumbers.ListNode{v, nil}
		tempNode1 = tempNode1.Next
	}


	return l1.Next

}

func validateAddTwoNumberstestData(node *addTwoNumbers.ListNode) ([]int) {
	data := []int{}
	for {
		data = append(data, node.Val)
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}

	return data
}
