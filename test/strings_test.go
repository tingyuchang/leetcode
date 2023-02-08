package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Strings"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	testData := []struct {
		input    int
		expected string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
	}

	for _, td := range testData {
		result := Strings.ConvertToTitle(td.input)
		assert.Equal(t, result, td.expected)
	}
}
