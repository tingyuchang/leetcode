package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestChuck(t *testing.T) {
	testData := []struct {
		input []int
		size int
		expected [][]int
	}{
		{input: []int{1,2,3,4}, size: 2,
			expected: [][]int{[]int{1,2}, []int{3,4}}},
		{input: []int{1,2,3,4,5}, size: 2,
			expected: [][]int{[]int{1,2}, []int{3,4}, []int{5}}},
		{input: []int{1,2,3,4,5,6,7,8}, size: 3,
			expected: [][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8}}},
		{input: []int{1,2,3,4,5}, size: 4,
			expected: [][]int{[]int{1,2,3,4}, []int{5}}},
		{input: []int{1,2,3,4,5}, size: 10,
			expected: [][]int{[]int{1,2,3,4,5}}},
	}

	for _,tt := range testData {
		result := chuck(tt.input, tt.size)
		assert.Equal(t, result, tt.expected)
	}
}