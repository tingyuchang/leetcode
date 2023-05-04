package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/basicProblems"
	"testing"
)

func TestRandomizedSelect(t *testing.T) {
	testData := []struct {
		nums   []int
		target int
		exp    int
	}{
		{[]int{6, 19, 4, 12, 14, 9, 15, 7, 8, 11, 3, 13, 2, 5, 10}, 5, 6},
	}

	for _, td := range testData {
		result := basicProblems.RandomizedSelect(td.nums, td.target)
		assert.Equal(t, result, td.exp)
	}
}
