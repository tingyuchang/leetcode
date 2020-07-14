package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/generateParenthesis"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testData := []struct {
		input    int
		expected []string
	}{
		{3,[]string{"()()()", "()(())", "(())()", "(()())", "((()))"}},
		{1,[]string{"()"}},
	}

	for _,td := range testData {
		result := generateParenthesis.GenerateParenthesis(td.input)
		if len(result) != len(td.expected) {
			t.Error("Result Not Equal")
			break
		}
		for _,v := range td.expected {
			for _,v2 := range result{
				assert.Equal(t, v2, v)
			}
		}
	}
}