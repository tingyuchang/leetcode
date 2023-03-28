package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Matrix"
	"testing"
)

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
