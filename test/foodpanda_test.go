package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/foodpanda"
	"testing"
)

func TestFoodpanda(t *testing.T) {
	testData := []struct {
		input    string
		expected string
	}{
		{"acb", "ab"},
		{"hot", "ho"},
		{"codility", "cdility"},
		{"aaaa", "aaa"},
	}

	for _, td := range testData {
		result := foodpanda.Solution2(td.input)
		assert.Equal(t, result, td.expected)
	}
}
