package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Sort"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	testData := []struct {
		n        int
		k        int
		expected string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
		{5, 37, "24135"},
	}

	for _, td := range testData {
		result := Sort.GetPermutation(td.n, td.k)
		assert.Equal(t, result, td.expected)
	}
}

func TestNextPermutation(t *testing.T) {
	testData := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 6, 5}},
		{[]int{1, 2, 3, 4, 6, 5}, []int{1, 2, 3, 5, 4, 6}},
		{[]int{1, 2, 3, 5, 4, 6}, []int{1, 2, 3, 5, 6, 4}},
	}

	for _, td := range testData {
		result := Sort.NextPermutation(td.input)
		assert.Equal(t, result, td.expected)
	}
}
