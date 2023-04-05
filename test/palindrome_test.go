package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		input    int
		expected bool
	}{
		{input: 121, expected: true},
		{input: -121, expected: false},
		{10, false},
	}

	for _, tt := range testData {
		result := Strings.IsPalindrome(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}
