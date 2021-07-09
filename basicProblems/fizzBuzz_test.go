package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testData := []struct{
		input int
		expected []string
	}{
		{input: 5, expected: []string{
			"1",
			"2",
			"fizz",
			"4",
			"buzz",
		}},
		{input: 18, expected: []string{
			"1",
			"2",
			"fizz",
			"4",
			"buzz",
			"fizz",
			"7",
			"8",
			"fizz",
			"buzz",
			"11",
			"fizz",
			"13",
			"14",
			"fizzbuzz",
			"16",
			"17",
			"fizz",
		}},
	}

	for _,tt := range testData {
		result := fizzBuzz(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}