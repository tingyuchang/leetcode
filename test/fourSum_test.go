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
	}

	for _,td := range testData {
		result := fourSum.FourSum(td.input, td.target)
		for i,v := range result {
			sort.Ints(result[i])
			sort.Ints(v)
			assert.Equal(t, result[i], v)
		}
	}

}
