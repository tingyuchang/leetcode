package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/singleNumber"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	var testData = []struct {
		input1   []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, tt := range testData {
		result := singleNumber.SingleNumber(tt.input1)
		assert.Equal(t, result, tt.expected)
	}

}
