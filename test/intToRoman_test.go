package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/intToRoman"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testData := []struct{
		input int
		expected string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58,"LVIII"},
		{1994,"MCMXCIV"},
	}

	for _,tt := range testData {
		result := intToRoman.IntToRoman(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}