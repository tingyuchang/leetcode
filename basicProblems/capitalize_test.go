package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCapitalize(t *testing.T) {
	testData := []struct{
		input string
		expected string
	} {
		{input: "hi there, how is it going?", expected: "Hi There, How Is It Going?"},
		{input: "i love breakfast at bill miller bbq", expected: "I Love Breakfast At Bill Miller Bbq"},
	}

	for _,tt := range testData {
		result := capitalize2(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}
func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		capitalize("hi there, how is it going?")
	}
}

func BenchmarkCapitalize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		capitalize2("hi there, how is it going?")
	}
}