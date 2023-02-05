package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/mergeSortedArray"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	var testData = []struct {
		input1   []int
		input2   int
		input3   []int
		input4   int
		expected []int
	}{
		{[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3, []int{-1, 0, 0, 1, 2, 2, 3, 3, 3}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, tt := range testData {
		result := mergeSortedArray.Merge(tt.input1, tt.input2, tt.input3, tt.input4)
		assert.Equal(t, result, tt.expected)
	}
}
