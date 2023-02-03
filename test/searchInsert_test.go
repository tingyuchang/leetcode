package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/searchInsertPosition"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var testData = []struct {
		input    []int
		input2   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, tt := range testData {
		result := searchInsertPosition.SearchInsertV2(tt.input, tt.input2)
		assert.Equal(t, result, tt.expected)
	}
}
