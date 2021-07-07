package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/nextPermutation"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testData := []struct{
		input []int
		expected []int
	}{
		{input: []int{1,2,3}, expected: []int{1,3,2}},
		{input: []int{3,2,31}, expected: []int{1,2,3}},
		{input: []int{1,1,5}, expected: []int{1,5,1}},
		{input: []int{1}, expected: []int{1}},
	}

	for _,tt := range testData {
		result := nextPermutation.NextPermutation(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}