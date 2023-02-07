package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/validPalindrome"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	var testData = []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}

	for _, tt := range testData {
		result := validPalindrome.IsPalindrome(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}
