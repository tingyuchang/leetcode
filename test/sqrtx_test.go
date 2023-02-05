package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/sqrtx"
	"testing"
)

func TestSqrtx(t *testing.T) {
	var testData = []struct {
		input1   int
		expected int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{100, 10},
	}

	for _, tt := range testData {
		result := sqrtx.MySqrt(tt.input1)
		assert.Equal(t, result, tt.expected)
	}

}
