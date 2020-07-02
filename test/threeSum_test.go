package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/threeSum"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testData := []struct{
		input []int
		expected [][]int
	}{
		{[]int{-1,0,1,2,-1,-4}, [][]int{[]int{-1,-1,2}, []int{-1,0,1}}},
		{[]int{-1,0,1,0}, [][]int{[]int{-1,0,1}}},
	}

	for _,tt := range testData {
		result := threeSum.ThreeSum(tt.input)

		assert.Equal(t, result, tt.expected)
	}
}