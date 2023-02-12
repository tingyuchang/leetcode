package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Array"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testData := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
	}

	for _, td := range testData {
		result := Array.ContainsDuplicate(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	testData := []struct {
		input    []int
		nearby   int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, td := range testData {
		result := Array.ContainsNearbyDuplicate(td.input, td.nearby)
		assert.Equal(t, result, td.expected)
	}
}
