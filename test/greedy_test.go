package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/greedy"
	"testing"
)

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
