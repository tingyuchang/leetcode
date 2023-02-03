package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/lengthOfLastWord"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	var testData = []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, tt := range testData {
		result := lengthOfLastWord.LengthOfLastWord(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}
