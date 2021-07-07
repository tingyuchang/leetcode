package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/findSubstring"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	testData := []struct {
		input string
		word []string
		expected []int
	}{
		{
			input: "wordgoodgoodgoodbestword",
			word: []string{"word","good","best","word"},
			expected: []int{},
		},
		{
			input: "barfoothefoobarman",
			word: []string{"foo", "bar"},
			expected: []int{0,9},
		},
		{
			input: "barfoofoobarthefoobarman",
			word: []string{"bar","foo","the"},
			expected: []int{6,9,12},
		},
	}

	for _,tt := range testData {
		result := findSubstring.FindSubstring(tt.input, tt.word)
		assert.Equal(t, result, tt.expected)
	}
}