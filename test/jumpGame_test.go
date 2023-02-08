package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/jumpGame"
	"testing"
)

func TestJumpGame(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 3, 2}, 2},
		{[]int{4, 3, 2, 1}, 1},
		{[]int{2, 1, 1, 1, 1}, 3},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}, 2},
	}

	for _, td := range testData {
		result := jumpGame.Jump(td.input)
		assert.Equal(t, result, td.expected)
	}

}
