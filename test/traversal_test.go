package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/traversal"
	"testing"
)

func TestMakeConnected(t *testing.T) {
	testData := []struct {
		n           int
		connections [][]int
		expected    int
	}{
		{
			4,
			[][]int{{0, 1}, {0, 2}, {1, 2}},
			1,
		},
		{
			6,
			[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}},
			2,
		},
		{
			6,
			[][]int{{0, 1}, {0, 2}, {4, 5}, {1, 2}, {1, 3}},
			1,
		},
		{
			6,
			[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 3}},
			-1,
		},
		{
			12,
			[][]int{{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2}},
			4,
		},
	}

	for _, td := range testData {
		result := traversal.MakeConnected(td.n, td.connections)
		assert.Equal(t, result, td.expected)
	}
}

func TestMinScore(t *testing.T) {
	testData := []struct {
		n        int
		roads    [][]int
		expected int
	}{
		{4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}, 5},
		{4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}, 2},
		{14, [][]int{
			{2, 9, 2308},
			{2, 5, 2150},
			{12, 3, 4944},
			{13, 5, 5462},
			{2, 10, 2187},
			{2, 12, 8132},
			{2, 13, 3666},
			{4, 14, 3019},
			{2, 4, 6759},
			{2, 14, 9869},
			{1, 10, 8147},
			{3, 4, 7971},
			{9, 13, 8026},
			{5, 12, 9982},
			{10, 9, 6459},
		}, 2150},
	}

	for _, td := range testData {
		result := traversal.MinScore(td.n, td.roads)
		assert.Equal(t, result, td.expected)
	}
}

func TestPacificAtlantic(t *testing.T) {
	testData := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{
			[][]int{{1}},
			[][]int{{0, 0}},
		},
		{
			[][]int{
				{10, 10, 10},
				{10, 1, 10},
				{10, 10, 10},
			},
			[][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{1, 0},
				{1, 2},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},
	}

	for _, td := range testData {
		result := traversal.PacificAtlantic(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestMaxAreaOfIsland(t *testing.T) {
	testData := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}, 6},
		{[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
		}, 0},
	}

	for _, td := range testData {
		result := traversal.MaxAreaOfIsland(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestSurroundedRegions(t *testing.T) {
	testData := []struct {
		input    [][]byte
		expected [][]byte
	}{
		//{[][]byte{
		//	{'X', 'X', 'X', 'X'},
		//	{'X', 'O', 'O', 'X'},
		//	{'X', 'X', 'O', 'X'},
		//	{'X', 'O', 'X', 'X'},
		//}, [][]byte{
		//	{'X', 'X', 'X', 'X'},
		//	{'X', 'X', 'X', 'X'},
		//	{'X', 'X', 'X', 'X'},
		//	{'X', 'O', 'X', 'X'},
		//}},
		//{[][]byte{
		//	{'X'},
		//}, [][]byte{
		//	{'X'},
		//}},
		//{[][]byte{
		//	{'O', 'O', 'O', 'O', 'X', 'X'},
		//	{'O', 'O', 'O', 'O', 'O', 'O'},
		//	{'O', 'X', 'O', 'X', 'O', 'O'},
		//	{'O', 'X', 'O', 'O', 'X', 'O'},
		//	{'O', 'X', 'O', 'X', 'O', 'O'},
		//	{'O', 'X', 'O', 'O', 'O', 'O'},
		//}, [][]byte{
		//	{'O', 'O', 'O', 'O', 'X', 'X'},
		//	{'O', 'O', 'O', 'O', 'O', 'O'},
		//	{'O', 'X', 'O', 'X', 'O', 'O'},
		//	{'O', 'X', 'O', 'O', 'X', 'O'},
		//	{'O', 'X', 'O', 'X', 'O', 'O'},
		//	{'O', 'X', 'O', 'O', 'O', 'O'},
		//}},
		{[][]byte{
			{'O', 'O', 'O', 'O'},
			{'O', 'O', 'O', 'O'},
			{'O', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'X'},
			{'O', 'X', 'X', 'X'},
			{'O', 'X', 'O', 'O'},
		}, [][]byte{
			{'O', 'O', 'O', 'O'},
			{'O', 'O', 'O', 'O'},
			{'O', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'X'},
			{'O', 'X', 'X', 'X'},
			{'O', 'X', 'O', 'O'},
		}},
	}

	for _, td := range testData {
		traversal.SolveSurroundedRegions(td.input)
		assert.Equal(t, td.input, td.expected)
	}
}

func TestNumsOfIslands(t *testing.T) {
	testData := []struct {
		input    [][]byte
		expected int
	}{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, 3},
		{[][]byte{{'0'}}, 0},
	}

	for _, td := range testData {
		result := traversal.NumIslands(td.input)
		assert.Equal(t, result, td.expected)
	}
}
