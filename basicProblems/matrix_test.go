package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMatrix(t *testing.T) {
	testData := []struct{
		input int
		expected [][]int
	}{
		{input: 1, expected: [][]int{[]int{1}}},
		{input: 2, expected: [][]int{
			[]int{1, 2},
			[]int{4, 3},
		}},
		{input: 3, expected: [][]int{
			[]int{1, 2, 3},
			[]int{8, 9, 4},
			[]int{7, 6, 5},
		}},
		{input: 4, expected: [][]int{
			[]int{1, 2, 3, 4},
			[]int{12, 13, 14, 5},
			[]int{11, 16, 15, 6},
			[]int{10, 9, 8, 7},
		}},
		{input: 5, expected: [][]int{
			[]int{1, 2, 3, 4, 5},
			[]int{16, 17, 18, 19, 6},
			[]int{15, 24, 25, 20, 7},
			[]int{14, 23, 22, 21, 8},
			[]int{13, 12, 11, 10, 9},
		}},
	}

	for _,tt := range testData{
		result := matrix(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}