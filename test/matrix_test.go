package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Matrix"
	"testing"
)

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
