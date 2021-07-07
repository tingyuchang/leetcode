package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	testData := []struct{
		input string
		expected string
	}{
		{input: "abcd", expected: "dcba"},
	}

	for _,tt := range testData {
		result := reverseString(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString("abcd")
	}
}
func BenchmarkReverseString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString2("abcd")
	}
}