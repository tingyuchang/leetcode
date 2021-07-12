package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCheckVowels(t *testing.T) {
	testData := []struct{
		input string
		expected int
	} {
		{input: "aeiou", expected: 5},
		{input: "AEIOU", expected: 5},
		{input: "abcdefghijklmnopqrstuvwxyz", expected: 5},
		{input: "why", expected: 0},
		{input: "bcdfghjkla", expected: 1},
	}

	for _,tt := range testData {
		result := checkVowels(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}