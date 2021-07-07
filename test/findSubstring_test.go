package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/findSubstring"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	testData := []struct {
		input    string
		words    []string
		expected []int
	}{
		{
			input:    "wordgoodgoodgoodbestword",
			words:    []string{"words","good","best","words"},
			expected: []int{},
		},
		{
			input:    "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0,9},
		},
		{
			input:    "barfoofoobarthefoobarman",
			words:    []string{"bar","foo","the"},
			expected: []int{6,9,12},
		},
	}

	for _,tt := range testData {
		result := findSubstring.FindSubstring(tt.input, tt.words)
		assert.Equal(t, result, tt.expected)
	}
}