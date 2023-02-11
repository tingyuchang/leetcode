package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/shortestAlternatingPaths"
	"testing"
)

func TestShortestAlternatingPaths(t *testing.T) {
	testData := []struct {
		n        int
		redEdge  [][]int
		blueEdge [][]int
		expected []int
	}{
		{3, [][]int{{0, 1}, {1, 2}}, [][]int{}, []int{0, 1, -1}},
		{3, [][]int{{1, 0}, {2, 1}}, [][]int{}, []int{0, -1, -1}},
		{3, [][]int{{0, 1}, {0, 2}}, [][]int{{1, 0}}, []int{0, 1, 1}},
		{3, [][]int{{0, 1}}, [][]int{{2, 1}}, []int{0, 1, -1}},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}, []int{0, 1, 2, 3, 7}},
	}

	for _, td := range testData {
		result := shortestAlternatingPaths.ShortestAlternatingPaths(td.n, td.redEdge, td.blueEdge)
		assert.Equal(t, result, td.expected)
	}
}
