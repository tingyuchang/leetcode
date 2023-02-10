package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Math"
	"testing"
)

func TestIsHappy(t *testing.T) {
	testData := []struct {
		input    int
		expected bool
	}{
		{19, true},
		{2, false},
		{1111111, true},
		{11, false},
	}

	for _, td := range testData {
		result := Math.IsHappy(td.input)
		assert.Equal(t, result, td.expected)
	}
}
