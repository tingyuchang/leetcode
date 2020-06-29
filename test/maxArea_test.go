package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/maxArea"
	"testing"
)

func TestMaxArea(t *testing.T) {
	var testData = []struct{
		intput []int
		expected int
	}{
		{intput: []int{1,8,6,2,5,4,8,3,7}, expected: 49},
	}

	for _,tt := range testData {
		result := maxArea.MaxArea(tt.intput)
		assert.Equal(t, result, tt.expected)
	}
}
