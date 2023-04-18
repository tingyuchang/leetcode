package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Matrix"
	"testing"
)

func TestFindBall(t *testing.T) {
	testData := []struct {
		grid [][]int
		exp  []int
	}{
		{
			[][]int{
				{1, 1, 1, -1, -1},
				{1, 1, 1, -1, -1},
				{-1, -1, -1, 1, 1},
				{1, 1, 1, 1, -1},
				{1, -1, -1, -1, -1},
			},
			[]int{
				1, -1, -1, -1, -1,
			},
		},
		{
			[][]int{
				{-1},
			},
			[]int{-1},
		},
		{
			[][]int{
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
			},
			[]int{
				0, 1, 2, 3, 4, -1,
			},
		},
		{
			[][]int{
				{-1, 1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, -1, -1, -1, 1, 1, 1, -1, -1, 1, 1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, 1, -1, 1, -1, -1, 1, 1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, -1, 1, 1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, 1, -1, -1, 1, 1, -1, -1, 1, -1, 1, -1, 1, 1, 1, -1, -1, -1, -1},
			},
			[]int{
				-1, -1, -1, 2, 3, 4, 5, 6, -1, -1, 9, 10, 11, 14, -1, -1, 15, 16, 19, 20, -1, -1, 21, 24, -1, -1, 25, -1, -1, 28, 29, 30, 31, 32, 33, 34, 35, -1, -1, -1, -1, 40, 41, 42, 43, 44, 45, -1, -1, 48, -1, -1, -1, -1, 53, 56, -1, -1, -1, -1, 59, 60, 61, 64, 65, 66, 67, 68, -1, -1, 71, 72, -1, -1, 75, 76, -1, -1, 77, 78, -1, -1, -1, -1, 83, 86, -1, -1, 87, -1, -1, -1, -1, 94, 95, -1, -1, 96, 97, 98,
			},
		},
	}

	for _, td := range testData {
		result := Matrix.FindBall(td.grid)
		assert.Equal(t, result, td.exp)
	}
}

func TestUpdateMatrix(t *testing.T) {
	testData := []struct {
		mat      [][]int
		expected [][]int
	}{
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		}, {
			[][]int{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
				{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
				{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
				{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
				{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
				{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
				{1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
			},
			[][]int{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
				{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
				{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
				{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
				{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
				{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{1, 0, 0, 0, 1, 2, 1, 1, 0, 1},
				{2, 1, 1, 1, 1, 2, 1, 0, 1, 0},
				{3, 2, 2, 1, 0, 1, 0, 0, 1, 1},
			},
		},
		{
			[][]int{
				{1, 1, 0, 0, 1, 0, 0, 1, 1, 0},
				{1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
				{1, 1, 1, 0, 0, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
				{0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
				{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
				{0, 1, 0, 1, 1, 0, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 0, 1, 1, 1, 1},
			},
			[][]int{
				{2, 1, 0, 0, 1, 0, 0, 1, 1, 0},
				{1, 0, 0, 1, 0, 1, 1, 2, 2, 1},
				{1, 1, 1, 0, 0, 1, 2, 2, 1, 0},
				{0, 1, 2, 1, 0, 1, 2, 3, 2, 1},
				{0, 0, 1, 2, 1, 2, 1, 2, 1, 0},
				{1, 1, 2, 3, 2, 1, 0, 1, 1, 1},
				{0, 1, 2, 3, 2, 1, 1, 0, 0, 1},
				{1, 2, 1, 2, 1, 0, 0, 1, 1, 2},
				{0, 1, 0, 1, 1, 0, 1, 2, 2, 3},
				{1, 2, 1, 0, 1, 0, 1, 2, 3, 4},
			},
		},
	}

	for _, td := range testData {
		result := Matrix.UpdateMatrix(td.mat)
		assert.Equal(t, result, td.expected)
	}
}

func TestRotateImage(t *testing.T) {
	testData := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			[][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, td := range testData {
		Matrix.RotateImage(td.matrix)
		assert.Equal(t, td.matrix, td.expected)
	}
}

func TestFindRatation(t *testing.T) {
	testData := []struct {
		mat      [][]int
		target   [][]int
		expected bool
	}{
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			[][]int{
				{1, 0},
				{0, 1},
			},
			true,
		},
		{
			[][]int{
				{0, 1},
				{1, 1},
			},
			[][]int{
				{1, 0},
				{0, 1},
			},
			false,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			[][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			true,
		},
	}

	for _, td := range testData {
		result := Matrix.FindRotation(td.mat, td.target)
		assert.Equal(t, result, td.expected)
	}
}

func TestSearchMatrix(t *testing.T) {
	testData := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1},
			},
			0,
			false,
		},
	}

	for _, td := range testData {
		result := Matrix.SearchMatrix(td.matrix, td.target)
		assert.Equal(t, result, td.expected)
	}
}

func TestSetZeroes(t *testing.T) {
	testData := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},

		{
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, td := range testData {
		Matrix.SetZeroes(td.matrix)
		assert.Equal(t, td.matrix, td.expected)
	}

}

func TestGenerateMatrix(t *testing.T) {
	testData := []struct {
		n        int
		expected [][]int
	}{
		{3, [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}},
		{
			1,
			[][]int{
				{1},
			},
		},
	}

	for _, td := range testData {
		result := Matrix.GenerateMatrix(td.n)
		assert.Equal(t, result, td.expected)
	}
}

func TestSpiralMatrix(t *testing.T) {
	testData := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, td := range testData {
		result := Matrix.SpiralOrder(td.matrix)
		assert.Equal(t, result, td.expected)
	}
}

func TestMaxDistance(t *testing.T) {
	testData := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, 2},
		{[][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 4},
		{[][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, -1},
	}
	for _, td := range testData {
		result := Matrix.MaxDistanceV2(td.input)
		assert.Equal(t, result, td.expected)
	}
}
