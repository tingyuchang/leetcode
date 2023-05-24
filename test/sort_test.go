package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Sort"
	"testing"
)

func TestMaxScore(t *testing.T) {
	testData := []struct {
		nums1 []int
		nums2 []int
		k     int
		exp   int64
	}{
		{
			[]int{1, 3, 3, 2},
			[]int{2, 1, 3, 4},
			3, 12,
		},
		{
			[]int{4, 2, 3, 1, 1},
			[]int{7, 5, 10, 9, 6},
			1, 30,
		},
	}

	for _, td := range testData {
		result := Sort.MaxScore(td.nums1, td.nums2, td.k)
		assert.Equal(t, result, td.exp)
	}
}

func TestLetterCasePermutation(t *testing.T) {
	testData := []struct {
		s        string
		expected []string
	}{
		{
			"a1b2",
			[]string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			"3z4",
			[]string{"3z4", "3Z4"},
		},
		{
			"C",
			[]string{"c", "C"},
		},
	}

	for _, td := range testData {
		result := Sort.LetterCasePermutation(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestKClosest(t *testing.T) {
	testData := []struct {
		pointer  [][]int
		k        int
		expected [][]int
	}{
		{
			[][]int{
				{1, 3},
				{-2, 2},
			},
			1,
			[][]int{
				{-2, 2},
			},
		},
		{
			[][]int{
				{3, 3},
				{5, -1},
				{-2, 4},
			},
			2,
			[][]int{
				{3, 3},
				{-2, 4},
			},
		},
	}

	for _, td := range testData {
		result := Sort.KClosest(td.pointer, td.k)
		assert.Equal(t, result, td.expected)
	}
}

func TestSuccessfulPairs(t *testing.T) {
	testData := []struct {
		spells   []int
		potions  []int
		success  int64
		expected []int
	}{
		{
			[]int{5, 1, 3},
			[]int{1, 2, 3, 4, 5},
			7,
			[]int{4, 0, 3},
		},
		{
			[]int{3, 1, 2},
			[]int{8, 5, 8},
			16,
			[]int{2, 0, 2},
		},
		{
			[]int{5, 23, 34, 2, 21, 14, 30, 23, 23, 14, 26, 3, 11, 7, 8, 13, 20, 12, 21, 35, 31, 8, 28, 32, 9, 26, 21, 20, 40, 24, 23, 32, 9, 18, 17, 23, 39, 38, 3, 8, 22, 26, 4, 10, 3, 15, 28, 11, 4, 5, 13, 16, 38, 6, 31, 37, 31, 31, 35, 40, 40, 5, 27, 11, 29, 13, 12, 14, 15, 14, 34, 19, 38, 26, 35, 12, 30, 3, 17, 7, 13, 8, 36, 14, 38, 26, 13, 34, 33, 33, 13, 23, 17},
			[]int{32, 39, 18, 34, 31, 32, 24, 36, 34, 24, 9, 36, 37, 16, 9, 20, 25, 26, 33, 29, 33, 18, 30, 37, 18, 38, 38, 38, 32, 12, 36, 18, 29, 5, 26, 4, 11, 38, 40, 15, 39, 36, 25, 35, 37, 5, 27, 23, 30, 24, 39, 9, 39, 22, 36, 21, 34, 24, 32, 32, 37, 17, 40, 36, 24, 32, 39, 40, 37, 19, 26, 40, 32, 34, 29, 26, 23, 25, 18, 5, 20, 26, 19, 40, 16, 34, 30, 29, 35, 40},
			195,
			[]int{11, 86, 86, 0, 83, 81, 86, 86, 86, 81, 86, 0, 77, 50, 59, 81, 83, 78, 83, 86, 86, 59, 86, 86, 67, 86, 83, 83, 89, 86, 86, 86, 67, 83, 82, 86, 89, 86, 0, 59, 86, 86, 0, 70, 0, 81, 86, 77, 0, 11, 81, 81, 86, 35, 86, 86, 86, 86, 86, 89, 89, 11, 86, 77, 86, 81, 78, 81, 81, 81, 86, 83, 86, 86, 86, 78, 86, 0, 82, 50, 81, 59, 86, 81, 86, 86, 81, 86, 86, 86, 81, 86, 82},
		},
		{
			[]int{1},
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, td := range testData {
		result := Sort.SuccessfulPairs(td.spells, td.potions, td.success)
		assert.Equal(t, result, td.expected)
	}

}

func TestRotateArray(t *testing.T) {
	testData := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			1,
			[]int{6, 1, 2, 3, 4, 5},
		},
	}

	for _, td := range testData {
		Sort.Rotate(td.nums, td.k)
		assert.Equal(t, td.nums, td.expected)
	}
}

func TestSortedSquares(t *testing.T) {
	testData := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{-1},
			[]int{1},
		},
		{
			[]int{-5, -3, -2, -1},
			[]int{1, 4, 9, 25},
		},
		{
			[]int{-10000, -9999, -7, -5, 0, 0, 10000},
			[]int{0, 0, 25, 49, 99980001, 100000000, 100000000},
		},
		{
			[]int{-10000, -1, 0, 3, 10000},
			[]int{0, 1, 9, 100000000, 100000000},
		},
		{
			[]int{-4, -4, -3},
			[]int{9, 16, 16},
		},
		{
			[]int{-2, 0},
			[]int{0, 4},
		},
	}

	for _, td := range testData {
		result := Sort.SortedSquares(td.nums)
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
