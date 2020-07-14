package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/findMedianSortedArrays"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testData := []struct{
		input []int
		input2 []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
		{[]int{}, []int{2,3}, 2.5},
		{[]int{7,7,7,7}, []int{2,3,9,10}, 7},
		{[]int{1,5}, []int{2,3,4,6}, 3.5},
	}

	for _,td := range testData {
		result := findMedianSortedArrays.FindMedianSortedArrays(td.input, td.input2)
		assert.Equal(t, result, td.expected)
	}
}