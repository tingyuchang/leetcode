package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSteps(t *testing.T) {
	testData := []struct{
		input int
		expected []string
	}{
		{input: 2, expected: []string{
			"# ",
			"##",
		}},
		{input: 3, expected: []string{
			"#  ",
			"## ",
			"###",
		}},
		{input: 4, expected: []string{
			"#   ",
			"##  ",
			"### ",
			"####",
		}},
	}

	for _,tt := range testData {
		result := steps(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}