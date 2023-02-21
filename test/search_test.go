package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Search"
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}

	for _, td := range testData {
		result := Search.SingleNonDuplicate(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestFindMinInRotatedArray(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 12, 13, 15}, 11},
		{[]int{1}, 1},
	}

	for _, td := range testData {
		result := Search.FindMinInRotatedArray(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestSearchRotatedArray(t *testing.T) {
	testData := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 9, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 2, -1},
		{[]int{3, 1}, 3, 0},
	}

	for _, td := range testData {
		result := Search.SearchRotatedArray(td.input, td.target)
		assert.Equal(t, result, td.expected)
	}
}
