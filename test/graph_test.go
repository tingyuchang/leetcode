package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/graph"
	"testing"
)

func TestDistanceLimitedPathsExist(t *testing.T) {
	testData := []struct {
		n        int
		edgeList [][]int
		queries  [][]int
		exp      []bool
	}{
		{
			3,
			[][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
			[][]int{{0, 1, 2}, {0, 2, 5}},
			[]bool{false, true},
		},
		{
			5,
			[][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}},
			[][]int{{0, 4, 14}, {1, 4, 13}},
			[]bool{true, false},
		},
	}

	for _, td := range testData {
		result := graph.DistanceLimitedPathsExist(td.n, td.edgeList, td.queries)
		assert.Equal(t, result, td.exp)
	}
}
