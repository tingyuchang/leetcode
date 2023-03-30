package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Sort"
	"testing"
)

func TestRemoveDuplicatesII(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, td := range testData {
		result := Sort.RemoveDuplicatesII(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, td := range testData {
		result := Sort.RemoveDuplicates(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestSearchInRotatedII(t *testing.T) {
	testData := []struct {
		nums     []int
		target   int
		expected bool
	}{
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			0, true,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			2, true,
		},
		{
			[]int{1, 0, 1, 1, 1},
			0, true,
		},
		{
			[]int{3, 1},
			1, true,
		},
	}

	for _, td := range testData {
		result := Sort.SearchInRotatedII(td.nums, td.target)
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
