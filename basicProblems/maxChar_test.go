package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMaxChar(t *testing.T) {
	testData := []struct{
		input string
		expected string
	} {
		{input: "abcdefghijklmnaaaaa", expected: "a"},
		{input: "ab1c1d1e1f1g1", expected: "1"},
		{input: "a", expected: "a"},
	}

	for _,tt := range testData {
		result := maxChar(tt.input)
		assert.Equal(t, result, tt.expected)
	}
 }

 func BenchmarkMaxChar(b *testing.B) {
 	for i:=0;i<b.N;i++{
 		maxChar("abcdefghijklmnaaaaa")
	}
 }