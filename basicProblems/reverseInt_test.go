package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	testData := []struct{
		input int
		expected int
	}{
		{input: 15, expected: 51},
		{input: 981, expected: 189},
		{input: 500, expected: 5},
		{input: -15, expected: -51},
		{input: -90, expected: -9},
	}

	for _,tt := range testData {
		result := reverseInt(tt.input)
		assert.Equal(t, result, tt.expected)
	}
	for _,tt := range testData {
		result := reverseInt2(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}

func BenchmarkReverseInt(b *testing.B) {
	for i:=0; i < b.N; i++ {
		reverseInt(982374123)
	}
}
func BenchmarkReverseInt2(b *testing.B) {
	for i:=0; i < b.N; i++ {
		reverseInt2(982374123)
	}
}