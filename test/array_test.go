package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Array"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	testData := []struct {
		s   string
		t   string
		exp bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "a#c#", true},
		{"a#c", "b", false},
	}

	for _, td := range testData {
		restult := Array.BackspaceCompare(td.s, td.t)
		assert.Equal(t, restult, td.exp)
	}

}

func TestValidateStackSequences(t *testing.T) {
	testData := []struct {
		pushed   []int
		popped   []int
		expected bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
	}

	for _, td := range testData {
		result := Array.ValidateStackSequences(td.pushed, td.popped)
		assert.Equal(t, result, td.expected)
	}

}

func TestNextGreaterElementsIII(t *testing.T) {
	testData := []struct {
		n        int
		expected int
	}{
		{12, 21},
		{21, -1},
		{230241, 230412},
		{2147483486, -1},
	}

	for _, td := range testData {
		result := Array.NextGreaterElementIII(td.n)
		assert.Equal(t, result, td.expected)
	}
}

func TestNextGreaterElementsII(t *testing.T) {
	testData := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 1},
			[]int{2, -1, 2},
		},
		{
			[]int{1, 2, 3, 4, 3},
			[]int{2, 3, 4, -1, 4},
		},
	}

	for _, td := range testData {
		result := Array.NextGreaterElementsII(td.nums)
		assert.Equal(t, result, td.expected)
	}
}

func TestNextGreaterElement(t *testing.T) {
	testData := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 42},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	}

	for _, td := range testData {
		result := Array.NextGreaterElement(td.nums1, td.nums2)
		assert.Equal(t, result, td.expected)
	}
}

func TestContainsDuplicate(t *testing.T) {
	testData := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
	}

	for _, td := range testData {
		result := Array.ContainsDuplicate(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	testData := []struct {
		input    []int
		nearby   int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, td := range testData {
		result := Array.ContainsNearbyDuplicate(td.input, td.nearby)
		assert.Equal(t, result, td.expected)
	}
}
