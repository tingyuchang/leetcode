package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/removeElement"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testData := []struct{
		input []int
		target int
		expected int
	}{
		{[]int{3,2,2,3}, 3, 2},
		{[]int{0,1,2,2,3,0,4,2}, 2, 5},
		{[]int{}, 0, 0},
		{[]int{2}, 3, 1},
	}

	for _,td := range testData {
		result := removeElement.RemoveElement(td.input, td.target)
		assert.Equal(t, result, td.expected)
	}
}