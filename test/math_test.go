package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Math"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testData := []struct {
		input    []int
		excepted []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{0, 4, 0}, []int{0, 0, 0}},
	}

	for _, td := range testData {
		result := Math.ProductExceptSelf(td.input)
		assert.Equal(t, result, td.excepted)
	}
}

func TestIsHappy(t *testing.T) {
	testData := []struct {
		input    int
		expected bool
	}{
		{19, true},
		{2, false},
		{1111111, true},
		{11, false},
	}

	for _, td := range testData {
		result := Math.IsHappy(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestCountOdds(t *testing.T) {
	testData := []struct {
		low      int
		high     int
		expected int
	}{
		{3, 7, 3},
		{8, 10, 1},
	}

	for _, td := range testData {
		result := Math.CountOdds(td.low, td.high)
		assert.Equal(t, result, td.expected)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	testData := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{16, true},
		{3, false},
		{0, false},
		{-16, false},
	}

	for _, td := range testData {
		result := Math.IsPowerOfTwo(td.input)
		assert.Equal(t, result, td.expected)
	}
}
