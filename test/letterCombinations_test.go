package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/letterCombinations"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testData := []struct{
		input string
		expected []string
	}{
		{"", []string{}},
		{"23",[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _,td := range testData {
		result := letterCombinations.LetterCombinations(td.input)
		assert.Equal(t, result, td.expected)
	}
}