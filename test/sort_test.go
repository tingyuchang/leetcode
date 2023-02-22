package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Sort"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testData := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 6, 5}},
		{[]int{1, 2, 3, 4, 6, 5}, []int{1, 2, 3, 5, 4, 6}},
		{[]int{1, 2, 3, 5, 4, 6}, []int{1, 2, 3, 5, 6, 4}},
	}

	for _, td := range testData {
		result := Sort.NextPermutation(td.input)
		assert.Equal(t, result, td.expected)
	}
}
