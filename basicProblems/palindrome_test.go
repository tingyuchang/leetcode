package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	testData := []struct{
		input string
		expected bool
	}{
		{input: "abba", expected: true},
		{input: "aba", expected: true},
		{input: " aba", expected: false},
		{input: "1000000001", expected: true},
	}

	for _,tt := range testData {
		result := palindrome(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i:=0; i < b.N; i++ {
		palindrome("abcd")
	}
}