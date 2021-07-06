package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/strStr"
	"testing"
)

func TestStrStr(t *testing.T) {
	var testData = []struct{
		haystack string
		needle string
		expected int
	}{
		{haystack: "hello", needle: "ll", expected: 2},
		{haystack: "aaaaa", needle: "bba", expected: -1},
		{haystack: "", needle: "", expected: 0},
	}

	for _,tt := range testData {
		result := strStr.StrStr(tt.haystack, tt.needle)
		assert.Equal(t, result, tt.expected)
	}
}