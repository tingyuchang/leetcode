package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/traversal"
	"testing"
)

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
