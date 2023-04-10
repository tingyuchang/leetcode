package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Array"
	"testing"
)

func TestNextGreaterElementsII(t *testing.T) {
	testData := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 1},
			[]int{2, -1, 2},
		},
		{
			[]int{1, 2, 3, 4, 3},
			[]int{2, 3, 4, -1, 4},
		},
	}

	for _, td := range testData {
		result := Array.NextGreaterElementsII(td.nums)
		assert.Equal(t, result, td.expected)
	}
}

func TestNextGreaterElement(t *testing.T) {
	testData := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 42},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	}

	for _, td := range testData {
		result := Array.NextGreaterElement(td.nums1, td.nums2)
		assert.Equal(t, result, td.expected)
	}
}

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
