package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/reverseInteger"
	"testing"
)

func TestReverse(t *testing.T) {
	var testData = []struct{
		input 		int
		expected 	int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for _,tt := range testData {
		result := reverseInteger.Reverse(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}