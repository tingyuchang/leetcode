package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/greedy"
	"testing"
)

func TestMinimizeArrayValue(t *testing.T) {
	testData := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 7, 1, 6}, 5},
		{[]int{10, 1}, 10},
	}

	for _, td := range testData {
		result := greedy.MinimizeArrayValue(td.nums)
		assert.Equal(t, result, td.expected)
	}
}

func TestCanPlaceFlowers(t *testing.T) {
	testData := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 1, 0, 0}, 1, true},
	}

	for _, td := range testData {
		result := greedy.CanPlaceFlowers(td.flowerbed, td.n)
		assert.Equal(t, result, td.expected)
	}
}

func TestEraseOverlapIntervals(t *testing.T) {
	testData := []struct {
		intervals [][]int
		expected  int
	}{
		{
			[][]int{
				{1, 2}, {2, 3}, {3, 4}, {1, 3},
			},
			1,
		},
	}

	for _, td := range testData {
		result := greedy.EraseOverlapIntervals(td.intervals)
		assert.Equal(t, result, td.expected)
	}
}
