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

func TestCountOdds(t *testing.T) {
	testData := []struct {
		low      int
		high     int
		expected int
	}{
		{3, 7, 3},
		{8, 10, 1},
	}

	for _, td := range testData {
		result := Math.CountOdds(td.low, td.high)
		assert.Equal(t, result, td.expected)
	}
}
