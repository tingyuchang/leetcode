package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFib(t *testing.T) {
	testData := []struct{
		input int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{39, 63245986},
	}

	for _,tt := range testData {
		result := fib(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}

func TestFib2(t *testing.T) {
	testData := []struct{
		input int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{39, 63245986},
	}

	for _,tt := range testData {
		result := fib2(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(39)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(39)
	}
}