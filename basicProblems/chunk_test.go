package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestChunk(t *testing.T) {
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
		result := chunk(tt.input, tt.size)
		assert.Equal(t, result, tt.expected)
	}
}

func TestChunk2(t *testing.T) {
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
		result := chunk(tt.input, tt.size)
		assert.Equal(t, result, tt.expected)
	}
}

func BenchmarkChunk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chunk([]int{1,2,3,4,5,6,7,8}, 3)
	}
}

func BenchmarkChunk2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chunk2([]int{1,2,3,4,5,6,7,8}, 3)
	}
}