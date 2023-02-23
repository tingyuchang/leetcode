package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/traversal"
	"testing"
)

func TestMinimumFuelCost(t *testing.T) {
	testData := []struct {
		roads    [][]int
		seats    int
		expected int64
	}{
		{[][]int{{0, 1}, {0, 2}, {0, 3}}, 5, 3},
		{[][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2, 7},
		{[][]int{}, 1, 0},
	}

	for _, td := range testData {
		result := traversal.MinimumFuelCost(td.roads, td.seats)
		assert.Equal(t, result, td.expected)
	}
}
