package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/graph"
	"testing"
)

func TestMaxNumEdgesToRemove(t *testing.T) {
	testData := []struct {
		n     int
		edges [][]int
		exp   int
	}{
		{4, [][]int{
			{3, 1, 2},
			{3, 2, 3},
			{1, 1, 3},
			{1, 2, 4},
			{1, 1, 2},
			{2, 3, 4},
		}, 2},
		{4, [][]int{
			{3, 1, 2},
			{3, 2, 3},
			{1, 1, 4},
			{2, 1, 4},
		}, 0},
		{
			4,
			[][]int{
				{3, 2, 3},
				{1, 1, 2},
				{2, 3, 4},
			},
			-1,
		},
		{
			4,
			[][]int{
				{3, 1, 2},
				{3, 3, 4},
				{1, 1, 3},
				{2, 2, 4},
			},
			0,
		},
		{
			2,
			[][]int{
				{1, 1, 2},
				{2, 1, 2},
				{3, 1, 2},
			},
			2,
		},
	}

	for _, td := range testData {
		result := graph.MaxNumEdgesToRemove(td.n, td.edges)
		assert.Equal(t, result, td.exp)
	}
}

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
