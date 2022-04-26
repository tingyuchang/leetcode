package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPyramid(t *testing.T) {
	testData := []struct {
		input int
		expected []string
	} {
		{1, []string{"#"}},
		{2,
			[]string{
			" # ",
			"###",
			},
		},
		{3,
			[]string{
				"  #  ",
				" ### ",
				"#####",
			},
		},
	}

	for _,tt := range testData {
		result := pyramid(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}
