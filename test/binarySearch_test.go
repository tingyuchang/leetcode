package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/BinarySearch"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testData := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}

	for _, td := range testData {
		result := BinarySearch.SearchRange(td.input, td.target)
		assert.Equal(t, result, td.expected)
	}
}
