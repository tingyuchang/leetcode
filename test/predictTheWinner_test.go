package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/predictTheWinner"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	testData := []struct{
		input []int
		expected bool
	}{
		{[]int{1,5,2}, false},
		{[]int{1,5,233,7}, true},
	}

	for _,td := range testData {
		result := predictTheWinner.PredictTheWinner(td.input)
		assert.Equal(t, result, td.expected)
	}
}