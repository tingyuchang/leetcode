package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/namingACompany"
	"testing"
)

func TestNamingACompany(t *testing.T) {
	testData := []struct {
		input    []string
		expected int64
	}{
		{[]string{"coffee", "donuts", "time", "toffee"}, 6},
		{[]string{"lack", "back"}, 0},
		{[]string{"bzklqtbdr", "kaqvdlp", "r", "dk"}, 12},
		{[]string{"phhrrjjcm", "zjfkpps", "pm", "fnpduelfe", "mxtvjnq"}, 18},
	}

	for _, td := range testData {
		result := namingACompany.DistinctNames(td.input)
		assert.Equal(t, result, td.expected)
	}
}
