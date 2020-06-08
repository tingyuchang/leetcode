package test

import (
	"github.com/stretchr/testify/assert"
	"leetcode/lengthOfLongestSubstring"
	"testing"
)


func TestLengthOfLongestSubstring(t *testing.T) {
	var testData = []struct {
		input        	string
		expected 		int
	}{
		{"abcbbc", 3},
		{" ", 1},
		{"", 0},
		{"bbbb", 1},
		{"ababa", 2},
	}

	for _, tt := range testData {
		actual := lengthOfLongestSubstring.LengthOfLongestSubstring(tt.input)
		assert.Equal(t, actual, tt.expected)
	}
}
