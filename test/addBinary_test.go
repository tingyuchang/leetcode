package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/addBinary"
	"testing"
)

func TestAddBinary(t *testing.T) {
	var testData = []struct {
		input1   string
		input2   string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, tt := range testData {
		result := addBinary.AddBinary(tt.input1, tt.input2)
		assert.Equal(t, result, tt.expected)
	}
}