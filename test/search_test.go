package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Search"
	"testing"
)

func TestSearchRotatedArray(t *testing.T) {
	testData := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 9, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 2, -1},
		{[]int{3, 1}, 3, 0},
	}

	for _, td := range testData {
		result := Search.SearchRotatedArray(td.input, td.target)
		assert.Equal(t, result, td.expected)
	}
}
