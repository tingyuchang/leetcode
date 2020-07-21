package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/removeDuplicates"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testData := []struct{
		input []int
		expected int
	}{
		{[]int{1,1,2}, 2},
		{[]int{0,0,1,1,1,2,2,3,3,4}, 5},
	}

	for _,td := range testData {
		result := removeDuplicates.RemoveDuplicates(td.input)
		assert.Equal(t, result, td.expected)
	}
}