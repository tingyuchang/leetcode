package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/validParentheses"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	testData := []struct{
		input string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"", true},
		{"(}", false},
		{"([)]", false},
		{"{[]}", true},
		{"]}", false},
	}

	for _,td := range testData {
		result := validParentheses.IsValid(td.input)
		assert.Equal(t, result, td.expected, td.input)
	}
}