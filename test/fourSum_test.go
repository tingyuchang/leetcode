package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/fourSum"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	testData := []struct{
		input []int
		target int
		expected [][]int
	}{
		{[]int{1, 0, -1, 0 ,-2, 2},
			0,
			[][]int{[]int{-1, 0, 0, -1}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}},
		},
		{
			[]int{-3, -1, 0, 2, 4, 5},
			2,
			[][]int{[]int{-3, -1, 2, 4}},
		},
	}

	for _,td := range testData {
		result := fourSum.FourSum(td.input, td.target)
		assert.Equal(t, len(result), len(td.expected))

		if len(result) == len(td.expected) {
			for i,v := range td.expected {
				sort.Ints(v)
				sort.Ints(result[i])
				assert.Equal(t, result[i], v)
			}
		}
	}
}
