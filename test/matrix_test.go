package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Matrix"
	"testing"
)

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
