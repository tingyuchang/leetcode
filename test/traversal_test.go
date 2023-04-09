package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/traversal"
	"testing"
)

func TestOrangeRotting(t *testing.T) {
	testData := []struct {
		grid     [][]int
		expected int
	}{
		{
			[][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			4,
		},
		{
			[][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			-1,
		},
		{
			[][]int{
				{0, 2},
			},
			0,
		},
		{
			[][]int{
				{2},
				{2},
				{1},
				{0},
				{1},
				{1},
			},
			-1,
		},
	}

	for _, td := range testData {
		result := traversal.OrangesRotting(td.grid)
		assert.Equal(t, result, td.expected)
	}

}

func TestNumEnclaves(t *testing.T) {
	testData := []struct {
		grid     [][]int
		expected int
	}{
		{
			[][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			3,
		},
		{
			[][]int{
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
				{1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0},
				{0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0},
				{1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1},
				{0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0},
				{1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1},
				{0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0},
				{1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0},
				{1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1},
			},
			7,
		},
	}

	for _, td := range testData {
		result := traversal.NumEnclaves(td.grid)
		assert.Equal(t, result, td.expected)
	}

}

func TestLongestCycle(t *testing.T) {
	testData := []struct {
		edges    []int
		expected int
	}{
		{[]int{59, 83, 46, 18, 45, 52, -1, -1, 46, -1, 75, 86, 89, -1, -1, -1, -1, 7, -1, 34, -1, -1, -1, 34, 82, -1, 75, 30, 34, -1, 87, 7, 35, -1, -1, 54, 72, -1, -1, -1, 29, 56, -1, 55, 32, 44, 62, -1, 80, -1, -1, 15, 81, -1, 32, -1, -1, 53, 81, 40, 81, 72, 68, -1, -1, -1, 87, 73, -1, -1, 55, -1, -1, -1, -1, -1, 53, 89, 38, 25, 16, 4, 71, 7, 33, -1, 42, 34, 29, 33, 1, 23, -1}, 3},
		{[]int{49, 29, 24, 24, -1, -1, -1, 2, 23, -1, 44, 47, 52, 49, 9, 31, 40, 34, -1, 53, 33, -1, 2, -1, 18, 31, 0, 9, 47, 35, -1, -1, -1, -1, 4, 12, 14, 19, -1, 4, 4, 43, 25, 11, 54, -1, 25, 39, 17, -1, 22, 44, -1, 44, 29, 50, -1, -1}, -1},
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
	}

	for _, td := range testData {
		result := traversal.LongestCycle(td.edges)
		assert.Equal(t, result, td.expected)
	}
}

func TestCountPairs(t *testing.T) {
	testData := []struct {
		n        int
		edges    [][]int
		expected int64
	}{
		{
			3,
			[][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			0,
		},
		{
			7,
			[][]int{
				{0, 2},
				{0, 5},
				{2, 4},
				{1, 6},
				{5, 4},
			},
			14,
		},
	}

	for _, td := range testData {
		result := traversal.CountPairs(td.n, td.edges)
		assert.Equal(t, result, td.expected)
	}
}

func TestFindWords(t *testing.T) {
	testData := []struct {
		board    [][]byte
		words    []string
		expected []string
	}{
		{
			[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"eat", "oath"},
		},
		{
			[][]byte{
				{'o', 'a', 'b', 'n'},
				{'o', 't', 'a', 'e'},
				{'a', 'h', 'k', 'r'},
				{'a', 'f', 'l', 'v'},
			},
			[]string{"oa", "oaa"},
			[]string{"oa", "oaa"},
		},
		{
			[][]byte{
				{'a', 'b', 'c'},
				{'a', 'e', 'd'},
				{'a', 'f', 'g'},
			},
			//[]string{"eaabcdgfa"},
			[]string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"},
			[]string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"},
		},
		{
			[][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			[]string{"abcd"},
			[]string{},
		},
	}

	for _, td := range testData {
		result := traversal.FindWords(td.board, td.words)
		assert.Equal(t, result, td.expected)
	}
}

func TestMinReOrder(t *testing.T) {
	testData := []struct {
		n          int
		connection [][]int
		expected   int
	}{
		{6, [][]int{
			{0, 1},
			{1, 3},
			{2, 3},
			{4, 0},
			{4, 5},
		}, 3},
		{5, [][]int{
			{1, 0},
			{1, 2},
			{3, 2},
			{3, 4},
		}, 2},
		{3, [][]int{
			{1, 0},
			{2, 0},
		}, 0},
	}

	for _, td := range testData {
		result := traversal.MinReorder(td.n, td.connection)
		assert.Equal(t, result, td.expected)
	}
}

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
