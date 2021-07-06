package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/divide"
	"testing"
)

func TestDivide(t *testing.T) {
	testData := []struct{
		dividend int
		divisor int
		expected int
	} {
		{dividend: 10, divisor: 3, expected: 3},
		{dividend: 7, divisor: -3, expected: -2},
		{dividend: 0, divisor: 1, expected: 0},
		{dividend: 1, divisor: 1, expected: 1},
		{dividend: -2147483648, divisor: -1, expected: 2147483647},
		{dividend: -2147483648, divisor: 1, expected: -2147483648},
	}

	for _,tt := range testData {
		result := divide.Divide(tt.dividend, tt.divisor)
		assert.Equal(t, result, tt.expected)
	}

}