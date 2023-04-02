package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Sort"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	testData := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		//{
		//	[]int{-7, -3, 2, 3, 11},
		//	[]int{4, 9, 9, 49, 121},
		//},
		//{
		//	[]int{1},
		//	[]int{1},
		//},
		//{
		//	[]int{-1},
		//	[]int{1},
		//},
		//{
		//	[]int{-5, -3, -2, -1},
		//	[]int{1, 4, 9, 25},
		//},
		//{
		//	[]int{-10000, -9999, -7, -5, 0, 0, 10000},
		//	[]int{0, 0, 25, 49, 99980001, 100000000, 100000000},
		//},
		//{
		//	[]int{-10000, -1, 0, 3, 10000},
		//	[]int{0, 1, 9, 100000000, 100000000},
		//},
		//{
		//	[]int{-4, -4, -3},
		//	[]int{9, 16, 16},
		//},
		//{
		//	[]int{-2, 0},
		//	[]int{0, 4},
		//},
	}

	for _, td := range testData {
		result := Sort.SortedSquares(td.nums)
		assert.Equal(t, result, td.expected)
	}
}

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
